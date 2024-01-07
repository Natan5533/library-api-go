package infra

import (
	"github.com/Natan5533/library-api-go/actors/api/handlers"
	"github.com/Natan5533/library-api-go/actors/database/repo"
	"github.com/Natan5533/library-api-go/core/services/authors"
	"github.com/Natan5533/library-api-go/core/services/libraries"
	"gorm.io/gorm"
)

type Container struct {
	LibraryRepo repo.LibraryRepo

	AuthorsRepo repo.AuthorRepo

	LibraryService libraries.Service

	LibraryHandler handlers.LibraryHandler

	AuthorsHandler handlers.AuthorHandler
}

func InitContainer(db *gorm.DB) *Container {
	// Init Repositories
	libraryRepo := repo.InitLibraryRepo(db)

	authorRepo := repo.InitAuthorRepo(db) // fetch author repo

	// Init Services
	libraryService := libraries.New(libraryRepo)

	authorService := authors.New(authorRepo) // fetch author service

	// Init Handlers
	libraryHandler := handlers.NewHandler(libraryService)

	authorHandler := handlers.NewAuthorsHandler(authorService)

	return &Container{
		LibraryRepo:    *libraryRepo,
		LibraryService: *libraryService,
		LibraryHandler: *libraryHandler,
		AuthorsHandler: *authorHandler,
	}
}
