package repo

import (
	"github.com/Natan5533/library-api-go/actors/database/models"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func InitBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (repo *BookRepo) Create(title string, authorId int) (int, error) {
	bookModel := models.NewBook(title, authorId)

	tx := repo.db.Create(bookModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return bookModel.ID, nil
}

func (repo *BookRepo) GetById(id int) (*models.Book, error) {
	var book models.Book
	result := repo.db.First(&book, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (repo *BookRepo) Delete(id int) error {
	var book models.Book

	_, err := repo.GetById(id)
	if err != nil {
		return err
	}

	result := repo.db.Delete(&book, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *BookRepo) Update(id int, bookParams *models.Book) error {
	var book models.Book
	tx := repo.db.First(&book, id)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Model(&book).Updates(bookParams)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
