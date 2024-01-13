package models

import "gorm.io/gorm"

type Library struct {
	gorm.Model

	ID      int      `gorm:"primaryKey;autoIncrement"`
	Address string   `gorm:"type:varchar(100);not null"`
	Name    string   `gorm:"type:varchar(100);not null"`
	Authors []Author `gorm:"foreignKey:LibraryID"`
}

func NewLibrary(name string, address string) *Library {
	return &Library{
		Address: address,
		Name:    name,
	}
}
