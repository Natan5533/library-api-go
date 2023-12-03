package repo

import (
	"testing"

	"github.com/Natan5533/library-api-go/actors/database"
	"github.com/Natan5533/library-api-go/actors/database/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB
var repo *LibraryRepo

func TestMain(m *testing.M) {
	db = database.Connect()
	repo = InitLibraryRepo(db)
	m.Run()
	TeardownDB(db, repo)
}

func TestInsert(t *testing.T) {
	t.Run("Success when we have a valid id", func(t *testing.T) {
		expectedName := "Kalunga"
		expectedAddress := "3 Street"

		id, err := repo.Create("Kalunga", "3 Street")
		if err != nil {
			panic(err)
		}
		library, err := repo.GetById(id)

		assert.Equal(t, expectedName, library.Name)
		assert.Equal(t, expectedAddress, library.Address)
		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Successful fetch", func(t *testing.T) {
		id, _ := repo.Create("Kalunga", "3 Street")
		address, err := repo.GetById(id)

		assert.Nil(t, err)
		assert.Equal(t, "Kalunga", address.Name)

	})
	t.Run("Return error when library not found", func(t *testing.T) {
		nonExistantId := 347878
		address, err := repo.GetById(nonExistantId)

		assert.Nil(t, address)
		assert.Equal(t, "record not found", err.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success deleting a library", func(t *testing.T) {
		expectedGetError := "record not found"
		libraryId, _ := repo.Create("Kalunga", "3 Street")
		err := repo.Delete(libraryId)
		_, errGet := repo.GetById(libraryId)

		assert.Nil(t, err)
		assert.Equal(t, expectedGetError, errGet.Error())
	})

	t.Run("Error deleting a library", func(t *testing.T) {
		nonExistantId := 347878
		err := repo.Delete(nonExistantId)
		assert.Equal(t, "a", err)
	})
}

func InitDB() (db *gorm.DB, repo *LibraryRepo) {
	a := database.Connect()
	return a, InitLibraryRepo(db)
}

func TeardownDB(db *gorm.DB, repo *LibraryRepo) {
	repo.db.Delete(models.Library{})
	db.Migrator().DropTable(models.Library{})
	sql, _ := db.DB()
	sql.Close()
}
