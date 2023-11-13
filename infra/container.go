package infra

import (
	"github.com/Natan5533/library-api-go/actors/api/handlers"
	"github.com/Natan5533/library-api-go/actors/database/repo"
	"github.com/Natan5533/library-api-go/core/services/librarysrv"
	"gorm.io/gorm"
)

type Container struct {
	LibraryRepo repo.LibraryRepo

	LibraryService librarysrv.Service

	LibraryHandler handlers.LibraryHandler
}

func InitContainer(db *gorm.DB) *Container {
	libraryRepo := repo.InitLibraryRepo(db)

	libraryService := librarysrv.New(libraryRepo)

	libraryHandler := handlers.NewHandler(libraryService)

	return &Container{
		LibraryRepo:    *libraryRepo,
		LibraryService: *libraryService,
		LibraryHandler: *libraryHandler,
	}
}
