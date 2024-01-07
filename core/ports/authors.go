package ports

import (
	"github.com/Natan5533/library-api-go/actors/database/models"
	adapters "github.com/Natan5533/library-api-go/core/domain/adpaters"
)

type AuthorsRepo interface {
	Create(email string, name string, libraryId int) (int, error)
	GetById(id int) (*models.Author, error)
}

type AuthorsService interface {
	Create(email string, name string, libraryId int) (int, error)
	GetById(id int) (*adapters.AuthorGetByIdResponse, error)
}
