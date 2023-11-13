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
		slog.Error("Deu ruim")
		return 0, tx.Error
	}

	return int(libraryModel.ID), nil
}

func (repo *LibraryRepo) GetById(id int) (*models.Library, error) {
	var library models.Library
	tx := repo.db.First(&library, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &library, nil
}
