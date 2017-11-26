package models

import "time"

const (
	StatusAccepted    = "accepted"
	StatusWrongAnswer = "wrong"
)

type Solution struct {
	ID         uint
	UserID     uint
	ExerciseID uint
	LessonID   uint
	CourseID   uint
	Status     string
	CreatedAt  time.Time
}
