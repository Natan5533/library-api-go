package repo

import (
	"testing"

	"github.com/Natan5533/library-api-go/actors/database/models"
	"github.com/stretchr/testify/assert"
)

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
		expectedGetError := "record not found"
		err := repo.Delete(nonExistantId)
		assert.Equal(t, expectedGetError, err.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success updating a library", func(t *testing.T) {
		libraryId, _ := repo.Create("Kalunga", "3 Street")
		library, _ := repo.GetById(libraryId)
		library.Name = "New Kalunga"
		err := repo.Update(libraryId, library)

		updatedLibrary, _ := repo.GetById(libraryId)

		assert.Nil(t, err)
		assert.Equal(t, "New Kalunga", updatedLibrary.Name)
	})

	t.Run("Error updating a library", func(t *testing.T) {
		nonExistantId := 347878
		expectedError := "record not found"

		err := repo.Update(nonExistantId, &models.Library{})

		assert.Equal(t, expectedError, err.Error())
	})
}
