package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	CookTime    string `json:"cookTime"`
	Ingredients datatypes.JSON
	Steps       datatypes.JSON
}

type User struct {
	gorm.Model
	Name     string `gorm:"size:255; not null"`
	Email    string `gorm:"uniqueIndex;size:255;not null"`
	Password string `gorm:"size:255 not null"`
}
