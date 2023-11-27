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
		expectedLibraryId := 1

		id, err := repo.Create("Test", "3 Street")
		library, _ := repo.GetById(id)

		assert.Equal(t, expectedLibraryId, int(library.ID))
		assert.Nil(t, err)
	})

	t.Run("Success when we have a valid id", func(t *testing.T) {
		expectedName := "Kalunga"
		expectedId := 2
		expectedAddress := "3 Street"

		id, err := repo.Create("Kalunga", "3 Street")
		if err != nil {
			panic(err)
		}
		library, err := repo.GetById(id)

		assert.Equal(t, expectedId, int(library.ID))
		assert.Equal(t, expectedName, library.Name)
		assert.Equal(t, expectedAddress, library.Address)
		assert.Nil(t, err)
	})

	t.Run("Error when Library not exists", func(t *testing.T) {
		expectedError := "record not found"
		library, err := repo.GetById(3)

		assert.Equal(t, expectedError, err.Error())
		assert.Nil(t, library)
	})

	t.Run("Success deleting a library", func(t *testing.T) {
		expectedGetError := "record not found"
		err := repo.Delete(1)
		_, errGet := repo.GetById(1)

		assert.Nil(t, err)
		assert.Equal(t, expectedGetError, errGet.Error())
	})

	t.Run("Error deleting a library", func(t *testing.T) {
		expectedError := "record not found"
		err := repo.Delete(3)

		assert.Equal(t, expectedError, err.Error())
	})

	TeardownDB(db)

}

func TeardownDB(db *gorm.DB) {
	db.Migrator().DropTable(models.Library{})
	sql, _ := db.DB()
	sql.Close()
}
