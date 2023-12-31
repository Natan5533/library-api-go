package infra

import (
	"github.com/Natan5533/library-api-go/actors/api/handlers"
	"github.com/Natan5533/library-api-go/actors/database/repo"
	"github.com/Natan5533/library-api-go/core/services/libraries"
	"gorm.io/gorm"
)

type Container struct {
	LibraryRepo repo.LibraryRepo

	LibraryService libraries.Service

	LibraryHandler handlers.LibraryHandler
}

func InitContainer(db *gorm.DB) *Container {
	libraryRepo := repo.InitLibraryRepo(db)

	libraryService := libraries.New(libraryRepo)

	libraryHandler := handlers.NewHandler(libraryService)

	return &Container{
		LibraryRepo:    *libraryRepo,
		LibraryService: *libraryService,
		LibraryHandler: *libraryHandler,
	}
}
