package database

import (
	"log/slog"

	"github.com/Natan5533/library-api-go/actors/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=library_api_go_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		slog.Error("[DB] Could not connect to database")
		panic(err)
	}

	if err := db.AutoMigrate(models.Library{}, models.Author{}, models.Book{}); err != nil {
		slog.Error("[DB] Error migrating")
		panic(err)
	}

	slog.Info("Successfully connected to database")

	return db
}
