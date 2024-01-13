package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model

	ID        int    `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"type:varchar(100);not null"`
	Name      string `gorm:"type:varchar(100);not null"`
	LibraryID int    `gorm:"not null"`
	Books     []Book `gorm:"foreignKey:AuthorId"`
}

func NewAuthor(email string, name string, libraryId int) *Author {
	return &Author{
		Email:     email,
		Name:      name,
		LibraryID: libraryId,
	}
}
