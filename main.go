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

	// library := models.NewLibrary("Kalunga", "Paulista Avenue")

	// library.Libraries.Create("Kalunga", "Paulista Avenue")

	// LibraryRepo.Create("Kalunga2", "Paulista Avenue2")

	// db.Create(&models.Library{
	// 	Name:    "Kalunga",
	// 	Address: "Paulista Avenue",
	// })

	// var library models.Library

	// tx := db.First(&library, 1)
	// if tx.Error != nil {
	// 	slog.Error("Deu ruim")
	// 	panic(tx.Error)
	// }
}
