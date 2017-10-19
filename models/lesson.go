package models

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

func (r *RequestNewLesson) Add() error {
	return nil
}
