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

func (repo *LibraryRepo) Create(name string, address string) error {
	library := models.NewLibrary(name, address)
	tx := repo.db.Create(library)

	if tx.Error != nil {
		slog.Error("Deu ruim")
		return tx.Error
	}

	return nil
}
