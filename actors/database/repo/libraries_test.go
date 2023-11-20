package repo

import (
	"testing"

	"github.com/Natan5533/library-api-go/actors/database"
	"github.com/Natan5533/library-api-go/actors/database/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsert(t *testing.T) {
	db := database.Connect()
	repo := InitLibraryRepo(db)

	t.Run("Sucessfull insert", func(t *testing.T) {
		_, err := repo.Create("Test", "3 Street")
		assert.Nil(t, err)

		TeardownDB(db)
	})

	t.Run("Sucessfull insert", func(t *testing.T) {
		_, err := repo.Create("Test", "3 Street")
		assert.Nil(t, err)

		TeardownDB(db)
	})

}

func TeardownDB(db *gorm.DB) {
	db.Migrator().DropTable(models.Library{})
	sql, _ := db.DB()
	sql.Close()
}
