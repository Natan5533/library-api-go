package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model

	Id        int    `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"type:varchar(100);not null"`
	Name      string `gorm:"type:varchar(100);not null"`
	LibraryId int    `gorm:"not null"`
	Books     []Book `gorm:"foreignKey:AuthorId"`
}
