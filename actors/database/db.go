package database

import (
	"log/slog"

	"github.com/Natan5533/library-api-go/actors/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=library_api_go_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("[DB] Deu problema na conexão")
		panic(err)
	}

	if err := db.AutoMigrate(models.Library{}, models.Author{}, models.Book{}); err != nil {
		slog.Error("[DB] Deu problema na migração")
		panic(err)
	}

	slog.Info("Conectado com sucesso!")

	return db
}
