package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title    string `gorm:"type:varchar(100);not null"`
	AuthorId int    `gorm:"not null"`
	Author   Author
}
