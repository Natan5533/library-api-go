package adapters

import "github.com/Natan5533/library-api-go/actors/database/models"

type (
	BookGetByIdResponse struct {
		Id       int    `json:"id"`
		Title    string `json:"title"`
		AuthorID int    `json:"author_id"`
	}

	UpdateBookParams struct {
		Title string `json:"title"`
	}
)

func BuildBookGetByIdResponse(book *models.Book) *BookGetByIdResponse {
	return &BookGetByIdResponse{
		Id:       book.ID,
		Title:    book.Title,
		AuthorID: book.AuthorId,
	}
}
