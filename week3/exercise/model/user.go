package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `binding:"required,min=6,max=100"`
	Password string
	Token    string
}
