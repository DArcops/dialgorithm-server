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

func GenerateToken(message []byte) ([]byte, error) {
	return encrypt.Encrypt(message)
}
