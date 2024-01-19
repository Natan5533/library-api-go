package infra

import (
	"github.com/Natan5533/library-api-go/actors/api/handlers"
	"github.com/Natan5533/library-api-go/actors/database/repo"
	"github.com/Natan5533/library-api-go/core/services/authors"
	"github.com/Natan5533/library-api-go/core/services/books"
	"github.com/Natan5533/library-api-go/core/services/libraries"
	"gorm.io/gorm"
)

type Container struct {
	LibraryRepo repo.LibraryRepo

	AuthorsRepo repo.AuthorRepo

	BookRepo repo.BookRepo

	LibraryService libraries.Service

	LibraryHandler handlers.LibraryHandler

	AuthorsHandler handlers.AuthorHandler

	BookHandler handlers.BooksHandler
}

func InitContainer(db *gorm.DB) *Container {
	// Init Repositories
	libraryRepo := repo.InitLibraryRepo(db)

	authorRepo := repo.InitAuthorRepo(db)

	bookRepo := repo.InitBookRepo(db)

	// Init Services
	libraryService := libraries.New(libraryRepo)

	authorService := authors.New(authorRepo)

	booksService := books.New(bookRepo)

	// Init Handlers
	libraryHandler := handlers.NewHandler(libraryService)

	authorHandler := handlers.NewAuthorsHandler(authorService)

	bookHandler := handlers.NewBooksHandler(booksService)

	return &Container{
		LibraryHandler: *libraryHandler,
		AuthorsHandler: *authorHandler,
		BookHandler:    *bookHandler,
	}
}
