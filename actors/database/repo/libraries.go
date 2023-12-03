package repo

import (
	"github.com/Natan5533/library-api-go/actors/database/models"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type LibraryRepo struct {
	db *gorm.DB
}

func InitLibraryRepo(db *gorm.DB) *LibraryRepo {
	return &LibraryRepo{db: db}
}

func (repo *LibraryRepo) Create(name, address string) (int, error) {
	libraryModel := models.NewLibrary(name, address)

	tx := repo.db.Create(libraryModel)
	if tx.Error != nil {
		slog.Error("[REPO] Error creating a Library")
		return 0, tx.Error
	}

	return int(libraryModel.ID), nil
}

func (repo *LibraryRepo) GetById(id int) (*models.Library, error) {
	var library models.Library
	result := repo.db.First(&library, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &library, nil
}

func (repo *LibraryRepo) Delete(id int) error {
	var library models.Library
	err := repo.db.Delete(&library, id).Error
	if err != nil {
		return err
	}

	return nil
}
