package libraries

import (
	"testing"

	"github.com/Natan5533/library-api-go/actors/database"
	"github.com/Natan5533/library-api-go/actors/database/models"
	"github.com/Natan5533/library-api-go/actors/database/repo"
	adapters "github.com/Natan5533/library-api-go/core/domain/adpaters"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var service *Service

func TestMain(m *testing.M) {
	db := database.Connect()
	repository := repo.InitLibraryRepo(db)
	service = New(repository)
	m.Run()
	TeardownDB(db, repository)
}

func TestCreate(t *testing.T) {
	t.Run("Success creating a library", func(t *testing.T) {
		expectedName := "Kalunga"
		expectedAddress := "3 Street"

		id, err := service.Create("Kalunga", "3 Street")
		if err != nil {
			panic(err)
		}
		library, err := service.GetById(id)

		assert.Equal(t, expectedName, library.Name)
		assert.Equal(t, expectedAddress, library.Address)
		assert.Nil(t, err)
	})

	t.Run("Return error when name is empty", func(t *testing.T) {
		expectedError := "name is empty"
		_, err := service.Create("", "3 Street")

		assert.Equal(t, expectedError, err.Error())
	})

	t.Run("Return error when address is empty", func(t *testing.T) {
		expectedError := "address is empty"
		_, err := service.Create("Kalunga", "")

		assert.Equal(t, expectedError, err.Error())
	})
}

func TestGetById(t *testing.T) {
	t.Run("Successful fetch", func(t *testing.T) {
		id, _ := service.Create("Kalunga", "3 Street")
		address, err := service.GetById(id)

		assert.Nil(t, err)
		assert.Equal(t, id, address.Id)
		assert.Equal(t, "Kalunga", address.Name)
		assert.Equal(t, "3 Street", address.Address)
	})

	t.Run("Return error when library not found", func(t *testing.T) {
		nonExistantId := 347878
		address, err := service.GetById(nonExistantId)

		assert.Nil(t, address)
		assert.Equal(t, "record not found", err.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Successful update", func(t *testing.T) {
		id, _ := service.Create("Kalunga", "3 Street")

		UpdateParams := adapters.UpdateLibraryParams{
			Name: "Kalunga 2",
		}

		err := service.Update(id, &UpdateParams)

		library, _ := service.GetById(id)

		assert.Equal(t, "Kalunga 2", library.Name)
		assert.Equal(t, "3 Street", library.Address)

		assert.Nil(t, err)

	})

	t.Run("Return error when library not found", func(t *testing.T) {
		nonExistantId := 347878

		err := service.Update(nonExistantId, &adapters.UpdateLibraryParams{})

		assert.Equal(t, "record not found", err.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success deleting a library", func(t *testing.T) {
		expectedGetError := "record not found"
		libraryId, _ := service.Create("Kalunga", "3 Street")
		err := service.Delete(libraryId)
		_, errGet := service.GetById(libraryId)

		assert.Nil(t, err)
		assert.Equal(t, expectedGetError, errGet.Error())
	})

	t.Run("Error deleting a library", func(t *testing.T) {
		nonExistantId := 347878
		expectedGetError := "record not found"
		err := service.Delete(nonExistantId)
		assert.Equal(t, expectedGetError, err.Error())
	})
}

func TeardownDB(db *gorm.DB, repo *repo.LibraryRepo) {
	db.Migrator().DropTable(models.Library{})
	sql, _ := db.DB()
	sql.Close()
}
