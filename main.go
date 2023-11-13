package main

import (
	"log/slog"

	"github.com/Natan5533/library-api-go/actors/api"
	"github.com/Natan5533/library-api-go/actors/database"
	"github.com/Natan5533/library-api-go/infra"
)

func main() {
	slog.Info("Hello World!")

	db := database.Connect()

	container := infra.InitContainer(db)

	api.InitServer(container)
}
