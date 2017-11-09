package models

type Level struct {
	ID            uint `json:"id"`
	CourseID      uint `json:"course_id" binding:"required"`
	Number        uint `json:"number"`
	BaseDirectory string
}

func GetLevels(courseID int) ([]Level, error) {
	levels := []Level{}
	return levels, db.Find(&levels, "course_id = ?", courseID).Error
}

func (l Level) GetLessons() ([]Lesson, error) {
	lessons := []Lesson{}
	return lessons, db.Find(&lessons, "level_id = ?", l.ID).Error
}
