package books

import (
	"github.com/Natan5533/library-api-go/core/domain/adapters"
	"github.com/Natan5533/library-api-go/core/ports"
)

type Service struct {
	bookRepo ports.BooksRepo
}

func New(bookRepo ports.BooksRepo) *Service {
	return &Service{bookRepo: bookRepo}
}

func (service *Service) Create(title string, authorId int) (int, error) {
	id, err := service.bookRepo.Create(title, authorId)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (service *Service) GetById(id int) (*adapters.BookGetByIdResponse, error) {
	book, err := service.bookRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return adapters.BuildBookGetByIdResponse(book), nil
}

func (service *Service) Delete(id int) error {
	err := service.bookRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) Update(id int, params *adapters.UpdateBookParams) error {
	book, err := service.bookRepo.GetById(id)
	if err != nil {
		return err
	}

	if params.Title != "" {
		book.Title = params.Title
	}

	service.bookRepo.Update(id, book)
	return nil
}
