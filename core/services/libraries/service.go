package libraries

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
		Authors: library.Authors,
	}, nil
}

func (service *Service) Delete(id int) error {
	err := service.libraryRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) Update(id int, params *adapters.UpdateLibraryParams) error {
	library, err := service.libraryRepo.GetById(id)
	if err != nil {
		return err
	}

	if params.Name != "" {
		library.Name = params.Name
	}

	if params.Address != "" {
		library.Address = params.Address
	}

	service.libraryRepo.Update(id, library)
	return nil
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
