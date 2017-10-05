package models

type User struct {
	Name     string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"pass" binding:"required"`
	//Permissions []*Permission  `gorm:"many2many:accesses;"`
}
