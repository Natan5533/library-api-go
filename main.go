package main

import (
	"log/slog"

	"github.com/Natan5533/library-api-go/actors/database"
	"github.com/Natan5533/library-api-go/actors/database/repo"
)

func main() {
	slog.Info("Hello World!")

	db := database.Connect()

	LibraryRepo := repo.InitLibraryRepo(db)

	LibraryRepo.Create("Kalunga", "Paulista Avenue")
}
