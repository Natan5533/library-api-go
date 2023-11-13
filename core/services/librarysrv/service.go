package librarysrv

import (
	"errors"

	adapters "github.com/Natan5533/library-api-go/core/domain/adpaters"
	"github.com/Natan5533/library-api-go/core/ports"
)

type Service struct {
	libraryRepo ports.LibrariesRepo
}

func New(libraryRepo ports.LibrariesRepo) *Service {
	return &Service{libraryRepo: libraryRepo}
}

func (service *Service) Create(name string, address string) (int, error) {
	err := checkCreateParams(name, address)
	if err != nil {
		return 0, err
	}
	id, err := service.libraryRepo.Create(name, address)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (service *Service) GetById(id int) (*adapters.LibraryGetByIdResponse, error) {
	library, err := service.libraryRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return &adapters.LibraryGetByIdResponse{
		Id:      int(library.ID),
		Name:    library.Name,
		Address: library.Address,
		Authors: []adapters.AuthorsResponse{},
	}, nil
}

func checkCreateParams(name string, address string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	if address == "" {
		return errors.New("address is empty")
	}
	return nil
}
