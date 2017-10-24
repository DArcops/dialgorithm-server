package models

type Subscription struct {
	ID       uint
	CourseID uint   `json:"course_id" binding:"required"`
	UserID   uint   `sql:"index"`
	UserPass string `gorm:"-" json:"user_pass" binding:"required"`
}

func (s *Subscription) Add() error {
	if db.First(&Course{}, "id = ?", s.CourseID).RecordNotFound() {
		return ErrNotFound
	}

	if !db.First(&Subscription{}, "course_id = ? and user_id = ?", s.CourseID, s.UserID).RecordNotFound() {
		return ErrDuplicate
	}

	subscription := Subscription{
		CourseID: s.CourseID,
		UserID:   s.UserID,
	}

	if err := db.Create(&subscription).Error; err != nil {
		return ErrToCreate
	}

	return nil
}
