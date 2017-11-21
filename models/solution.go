package models

import "time"

const (
	StatusAccepted    = "accpeted"
	StatusWrongAnswer = "wrong"
)

type Solution struct {
	ID         uint
	UserID     uint
	ExerciseID uint
	Status     string
	CreatedAt  time.Time
}
