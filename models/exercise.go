package models

import (
	"errors"
	"io/ioutil"
	"os"
)

type Exercise struct {
	ID            uint   `json:"id"`
	LessonID      uint   `json:"-"`
	Name          string `json:"name"`
	BaseDirectory string `json:"-"`
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
