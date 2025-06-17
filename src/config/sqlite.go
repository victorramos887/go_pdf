package config

import (
	"os"

	"github.com/victorramos887/go_pdf/src/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializerSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")
	dbPath := "./src/db/main.db"
	dir := "./src/db"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	_, statErr := os.Stat(dbPath)

	if os.IsNotExist(statErr) {
		logger.Info("database file not found, creating ...")
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Maintenance{})

	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil

}
