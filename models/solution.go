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

type CountStatusSolution struct {
	Accepted int `json:"accepted"`
	Wrong    int `json:"wrong"`
}

func GetSolutions() CountStatusSolution {
	ac := []Solution{}
	db.Find(&ac, "status = ?", StatusAccepted)
	wrg := []Solution{}
	db.Find(&wrg, "status = ?", StatusWrongAnswer)
	res := CountStatusSolution{
		Accepted: len(ac),
		Wrong:    len(wrg),
	}
	return res
}
