package authors

import (
	adapters "github.com/Natan5533/library-api-go/core/domain/adpaters"
	"github.com/Natan5533/library-api-go/core/ports"
)

type Service struct {
	authorRepo ports.AuthorsRepo
}

func New(authorRepo ports.AuthorsRepo) *Service {
	return &Service{authorRepo: authorRepo}
}

func (service *Service) Create(email string, name string, libraryId int) (int, error) {
	id, err := service.authorRepo.Create(email, name, libraryId)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (service *Service) GetById(id int) (*adapters.AuthorGetByIdResponse, error) {
	author, err := service.authorRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return &adapters.AuthorGetByIdResponse{
		Id:    int(author.ID),
		Email: author.Email,
		Name:  author.Name,
		Books: []adapters.Books{},
	}, nil
}
