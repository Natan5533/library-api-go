package test

import (
	"github.com/Natan5533/library-api-go/actors/database/models"
	"gorm.io/gorm"
)

// func TestSetup(m *testing.M) {
// 	// SetMain
// 	db := setup()
// 	m.Run()
// 	teardown(db)
// }

// func setup() *gorm.DB {
// 	return database.Connect()
// }

func TeardownDB(db *gorm.DB) {
	db.Delete(models.Library{}, models.Author{}, models.Book{})
}
