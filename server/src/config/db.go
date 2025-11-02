package config

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabaseConnection() *gorm.DB {
	var db, err = gorm.Open(postgres.Open(Config["Db"]), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect database", err)
	}

	return db
}
