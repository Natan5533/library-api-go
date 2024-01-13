package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	ID       int    `gorm:"primaryKey;autoIncrement"`
	Title    string `gorm:"type:varchar(100);not null"`
	AuthorId int    `gorm:"not null"`
}
