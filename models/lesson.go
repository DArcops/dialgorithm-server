package models

import (
	"errors"
	"io/ioutil"
	"os"
)

type Lesson struct {
	ID            uint   `json:"id"`
	LevelID       uint   `json:"-"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	BaseDirectory string `json:"-"`
	SummerNotCode string `json:"code" gorm:"-"`
}

type RequestNewLesson struct {
	Name        string `json:"name" binding:"required"`
	CourseID    uint   `json:"course_id" binding:"required"`
	LevelID     uint   `json:"level_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Markup      string `json:"code" binding:"required"`
}

func (r *RequestNewLesson) Add() error {

	level := Level{}
	if db.First(&level, "course_id = ? and id = ?", r.CourseID, r.LevelID).RecordNotFound() {
		return errors.New("record not found")
	}

	lesson := Lesson{
		Name:        r.Name,
		LevelID:     level.ID,
		Description: r.Description,
	}

	tx := db.Begin()
	if err := db.Create(&lesson).Error; err != nil {
		tx.Rollback()
		return err
	}

	lessonPath := level.BaseDirectory + "/" + "Lesson_" + lesson.Name
	if err := os.Mkdir(lessonPath, os.FileMode(0777)); err != nil {
		tx.Rollback()
		return err
	}

	if err := db.Model(lesson).Update("base_directory", lessonPath).Error; err != nil {
		tx.Rollback()
		return err
	}

	code := r.Markup
	err := ioutil.WriteFile(lessonPath+"/overview.html", []byte(code), 0777)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

func GetLessons(levelID uint, last int64) ([]Lesson, error) {
	lessons := []Lesson{}
	if last != -1 {
		return lessons, db.Find(&lessons, "level_id = ? and id > ?", levelID, last).Limit(4).Error
	} else {
		return lessons, db.Find(&lessons, "level_id = ?", levelID).Error
	}
}

func (l *Lesson) FillCode() error {
	data, err := ioutil.ReadFile(l.BaseDirectory + "/overview.html")
	if err != nil {
		return ErrToCreate
	}
	l.SummerNotCode = string(data)
	return nil
}
