package repo

import (
	"log/slog"

	"github.com/Natan5533/library-api-go/actors/database/models"
	"gorm.io/gorm"
)

type AuthorRepo struct {
	db *gorm.DB
}

func InitAuthorRepo(db *gorm.DB) *AuthorRepo {
	return &AuthorRepo{db: db}
}

func (repo *AuthorRepo) Create(email, name string, libraryId int) (int, error) {
	authorModel := models.NewAuthor(email, name, libraryId)

	tx := repo.db.Create(authorModel)
	if tx.Error != nil {
		slog.Error("[REPO] Error creating a Author", tx.Error)

		return 0, tx.Error
	}

	return int(authorModel.ID), nil
}

func (repo *AuthorRepo) GetById(id int) (*models.Author, error) {
	var author models.Author
	result := repo.db.Preload("Books").First(&author, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &author, nil
}

func (repo *AuthorRepo) Delete(id int) error {
	var author models.Author

	_, err := repo.GetById(id)
	if err != nil {
		return err
	}

	result := repo.db.Delete(&author, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *AuthorRepo) Update(id int, authorParams *models.Author) error {
	var author models.Author
	tx := repo.db.First(&author, id)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Model(&author).Updates(authorParams)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
