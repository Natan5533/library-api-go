package ports

import (
	"github.com/Natan5533/library-api-go/actors/database/models"
	"github.com/Natan5533/library-api-go/core/domain/adapters"
)

type (
	BooksRepo interface {
		Create(title string, authorId int) (int, error)
		GetById(id int) (*models.Book, error)
		Delete(id int) error
		Update(id int, bookParams *models.Book) error
	}

	BooksService interface {
		Create(title string, authorId int) (int, error)
		GetById(id int) (*adapters.BookGetByIdResponse, error)
		Delete(id int) error
		Update(id int, bookParams *adapters.UpdateBookParams) error
	}
)
