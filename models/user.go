package models

import "github.com/darcops/dialgorithm-server/modules/encrypt"

type User struct {
	ID       uint
	Name     string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"pass,omitempty" binding:"required"`
	CanWrite bool   `json:"administrator"`
	//Permissions []*Permission  `gorm:"many2many:accesses;"`
}

type NewAdmin struct {
	Email string `json:"email" binding:"required"`
}

func GenerateToken(message []byte) ([]byte, error) {
	return encrypt.Encrypt(message)
}

func (u User) GetCourses() ([]Course, error) {
	sub := []Subscription{}
	if err := db.Find(&sub, "user_id = ?", u.ID).Error; err != nil {
		return nil, err
	}
	courses := []Course{}
	for _, v := range sub {
		course := Course{}
		db.First(&course, "id = ?", v.CourseID)
		courses = append(courses, course)
	}
	return courses, nil
}

func (u User) AddNewAdmin(email string) error {
	return db.Model(&User{}).Where("email = ?", email).Update("can_write", true).Error
}
