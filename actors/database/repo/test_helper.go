package repo

import (
	"testing"

	"github.com/Natan5533/library-api-go/actors/database"
	"github.com/Natan5533/library-api-go/actors/database/models"
	"gorm.io/gorm"
)

var repo *LibraryRepo

// var authorRepo *AuthorRepo

func TestMain(m *testing.M) {
	db := database.Connect()
	TeardownDB(db)
	repo = InitLibraryRepo(db)
	m.Run()
	TeardownDB(db)
}

func TeardownDB(db *gorm.DB) {
	db.Migrator().DropTable(models.Library{})
	sql, _ := db.DB()
	sql.Close()
}
