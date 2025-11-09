package config

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func OpenDatabaseConnection() *gorm.DB {
	var db, err = gorm.Open(postgres.Open(Config.Db), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect database", "error", err)
	} else {
		slog.Info("Database connected.")
	}

	return db
}

func init() {
	slog.Info("Connecting to database...")
	Db = OpenDatabaseConnection()
}
