package utils

import (
	"fmt"
	"myapp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Establishes a connection to the database and returns a *gorm.DB object.
func DbConnection(dbname string) (*gorm.DB, error) {
	// Construct DSN string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword)

	// Add dbname to DSN if provided
	if dbname != "" {
		dsn += fmt.Sprintf(" dbname=%s", dbname)
	}

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db, nil
}
