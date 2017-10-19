package models

import (
	"errors"
	"io/ioutil"
	"os"
)

type Lesson struct {
	ID            uint
	LevelID       uint
	Name          string
	BaseDirectory string
}

type RequestNewLesson struct {
	Name     string `json:"name" binding:"required"`
	CourseID uint   `json:"course_id" binding:"required"`
	LevelID  uint   `json:"level_id" binding:"required"`
	Markup   string `json:"code" binding:"required"`
}

var htmlBase = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Summernote</title>
  <link href="http://netdna.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.8/summernote.css" rel="stylesheet">
  <script src="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.8/summernote.js"></script>
</head>
<body>`

//var err Error

func (r *RequestNewLesson) Add() error {

	level := Level{}
	if db.First(&level, "course_id = ? and id = ?", r.CourseID, r.LevelID).RecordNotFound() {
		return errors.New("record not found")
	}

	lesson := Lesson{
		Name:    r.Name,
		LevelID: level.ID,
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

	code := htmlBase + "\n" + r.Markup + "</body>"
	err := ioutil.WriteFile(lessonPath+"/overview.html", []byte(code), 0777)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}
