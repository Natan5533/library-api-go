package authors

import (
	"github.com/Natan5533/library-api-go/core/domain/adapters"
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
	return adapters.BuildAuthorGetByIdResponse(author), nil
}

func (service *Service) Delete(id int) error {
	err := service.authorRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) Update(id int, params *adapters.UpdateAuthorParams) error {
	author, err := service.authorRepo.GetById(id)
	if err != nil {
		return err
	}

	if params.Name != "" {
		author.Name = params.Name
	}

	if params.Email != "" {
		author.Email = params.Email
	}

	service.authorRepo.Update(id, author)
	return nil
}
