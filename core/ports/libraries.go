package ports

import "github.com/Natan5533/library-api-go/actors/database/models"

type Libraries interface {
	Create(name string, address string) (*models.Library, error)
	GetById(id uint) (*models.Library, error)
}
