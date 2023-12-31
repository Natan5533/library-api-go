package ports

import (
	"github.com/Natan5533/library-api-go/actors/database/models"
	adapters "github.com/Natan5533/library-api-go/core/domain/adpaters"
)

type LibrariesRepo interface {
	Create(name string, address string) (int, error)
	GetById(id int) (*models.Library, error)
	Delete(id int) error
	Update(id int, libraryParams *models.Library) error
}

type LibrariesService interface {
	Create(name string, address string) (int, error)
	GetById(id int) (*adapters.LibraryGetByIdResponse, error)
	Delete(id int) error
	Update(id int, libraryParams *adapters.UpdateLibraryParams) error
}
