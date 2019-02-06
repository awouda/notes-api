package domain

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Content string `json:"content"`
}

