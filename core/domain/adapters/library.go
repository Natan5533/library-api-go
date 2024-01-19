package adapters

import "github.com/Natan5533/library-api-go/actors/database/models"

type LibraryGetByIdResponse struct {
	Id      int               `json:"id"`
	Name    string            `json:"name"`
	Address string            `json:"address"`
	Authors []AuthorsResponse `json:"authors"`
}

type AuthorsResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateLibraryParams struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func BuildLibraryGetByIdResponse(library *models.Library) *LibraryGetByIdResponse {
	authors := make([]AuthorsResponse, len(library.Authors))
	for i, author := range library.Authors {
		authors[i] = AuthorsResponse{
			Id:   int(author.ID),
			Name: author.Name,
		}
	}

	return &LibraryGetByIdResponse{
		Id:      int(library.ID),
		Name:    library.Name,
		Address: library.Address,
		Authors: authors,
	}
}
