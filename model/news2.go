package model

import "gorm.io/gorm"

type News2 struct {
	gorm.Model
	Title   string `gorm:"type:varchar(256);not " json:"title"`
	Content string `gorm:"size:255;not null" json:"content"`
}
