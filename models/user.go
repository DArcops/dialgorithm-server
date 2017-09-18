package models

type User struct {
	Name     string
	Email    string
	Password string
	//Permissions []*Permission  `gorm:"many2many:accesses;"`
}
