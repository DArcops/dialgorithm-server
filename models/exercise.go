package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Exercise struct {
	ID            uint   `json:"id"`
	LessonID      uint   `json:"-"`
	Name          string `json:"name"`
	BaseDirectory string `json:"-"`
	Code          string `json:"code" gorm:"-"`
}

type RequestNewExercise struct {
	Name     string `json:"name" binding:"required"`
	CourseID uint   `json:"course_id" binding:"required"`
	LevelID  uint   `json:"level_id" binding:"required"`
	LessonID uint   `json:"lesson_id" binding:"required"`
	Markup   string `json:"code" binding:"required"`
	Input    string `json:"input"`
	Output   string `json:"output"`
}

var coliruUrl = "http://coliru.stacked-crooked.com/compile"
var cmdCompile = "g++ -std=c++17 -O2 -Wall -pedantic -pthread main.cpp && ./a.out"

type Coliru struct {
	Cmd string `json:"cmd"`
	Src string `json:"src"`
}

func (e *RequestNewExercise) Add() error {
	level := Level{}
	if db.First(&level, "course_id = ? and id = ?", e.CourseID, e.LevelID).RecordNotFound() {
		return errors.New("record not found")
	}
	lesson := Lesson{}
	if db.First(&lesson, "id = ? and level_id = ?", e.LessonID, level.ID).RecordNotFound() {
		return errors.New("record not found")
	}

	exercise := Exercise{
		Name:     e.Name,
		LessonID: e.LessonID,
	}

	tx := db.Begin()
	if err := db.Create(&exercise).Error; err != nil {
		tx.Rollback()
		return err
	}

	exercisePath := lesson.BaseDirectory + "/" + "Exercise_" + exercise.Name
	if err := os.Mkdir(exercisePath, os.FileMode(0777)); err != nil {
		tx.Rollback()
		return err
	}

	if err := db.Model(exercise).Update("base_directory", exercisePath).Error; err != nil {
		tx.Rollback()
		return err
	}

	code := e.Markup
	err := ioutil.WriteFile(exercisePath+"/instructions.html", []byte(code), 0777)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(e.Input) > 0 {
		err := ioutil.WriteFile(exercisePath+"/input.txt", []byte(e.Input), 0777)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(e.Output) > 0 {
		err := ioutil.WriteFile(exercisePath+"/output.txt", []byte(e.Output), 0777)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}

func (e Exercise) TestSolution(code string) string {
	dat, _ := ioutil.ReadFile(e.BaseDirectory + "/input.txt")

	strDat := string(dat)

	coliru := Coliru{
		Cmd: cmdCompile + " " + strDat,
		Src: code,
	}
	str, _ := json.Marshal(coliru)

	req, err := http.NewRequest("POST", coliruUrl, bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("COOODE STATUS", resp.StatusCode)
	return string(body)

}

func (e Exercise) Solve(code string, userID uint) (string, string) {
	compilation := e.TestSolution(code)

	dat, _ := ioutil.ReadFile(e.BaseDirectory + "/output.txt")

	strDat := string(dat)

	formatedSavedOutput := formatOutput(strDat)
	formatedSentSolution := formatOutput(compilation)

	fmt.Println("sent:", formatedSentSolution)
	fmt.Println("saved:", formatedSavedOutput)

	if formatedSentSolution == formatedSavedOutput {
		e.insertSolution(userID, StatusAccepted)
		return compilation, "Acepted"
	}
	e.insertSolution(userID, StatusWrongAnswer)
	return compilation, "Wrong"
}

func (e Exercise) insertSolution(userID uint, status string) {
	acceptedSolution := Solution{
		UserID:     userID,
		ExerciseID: e.ID,
		Status:     status,
	}
	db.Create(&acceptedSolution)
}

func formatOutput(strDat string) string {
	formatedSolution := ""
	for i := 0; i < len(strDat); i++ {
		if string(strDat[i]) != "\n" && string(strDat[i]) != "\r" && string(strDat[i]) != " " {
			formatedSolution += string(strDat[i])
		}
	}
	return formatedSolution
}
