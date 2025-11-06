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
		slog.Error("failed to connect database", "error", err)
	}

	return db
}

func init() {
	slog.Info("Connecting to database...")
	Db = OpenDatabaseConnection()
	slog.Info("Database connected.")
}
