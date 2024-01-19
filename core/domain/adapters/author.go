package adapters

import "github.com/Natan5533/library-api-go/actors/database/models"

type AuthorGetByIdResponse struct {
	Id    int             `json:"id"`
	Name  string          `json:"name"`
	Email string          `json:"email"`
	Books []BooksResponse `json:"books"`
}

type BooksResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type UpdateAuthorParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func BuildAuthorGetByIdResponse(author *models.Author) *AuthorGetByIdResponse {
	books := make([]BooksResponse, len(author.Books))
	for i, book := range author.Books {
		books[i] = BooksResponse{
			Id:    int(book.ID),
			Title: book.Title,
		}
	}
	return &AuthorGetByIdResponse{
		Id:    author.ID,
		Name:  author.Name,
		Email: author.Email,
		Books: books,
	}
}
