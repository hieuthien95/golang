package model

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title     string `binding:"required,min=6,max=100"`
	Completed bool
}
