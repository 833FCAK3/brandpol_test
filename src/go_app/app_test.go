package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/utils"
	"os"
	"testing"

	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	// Set up test database
	var err error
	testDB, err = createTestDatabase(config.Test_db)
	if err != nil {
		log.Fatalf("Error creating test database: %v\n", err)
	}

	// Run tests
	exitCode := m.Run()

	// Clean up test database
	if err := dropTestDatabase(config.Test_db); err != nil {
		log.Fatalf("Error dropping test database %v: %v\n", config.Test_db, err)
	}

	os.Exit(exitCode)
}

func createTestDatabase(dbname string) (*gorm.DB, error) {
	// Connect to the PostgreSQL server
	db, err := utils.DbConnection("")
	if err != nil {
		return nil, err
	}

	var count int64
	db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbname).Scan(&count)
	if count > 0 {
		// Database already exists, return without creating it
		return nil, nil
	}

	// Create the test database
	rawSQL := fmt.Sprintf("CREATE DATABASE %s", dbname)
	if err := db.Exec(rawSQL).Error; err != nil {
		return nil, err
	}

	// Connect to the newly created test database
	db, err = utils.DbConnection(dbname)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func dropTestDatabase(dbName string) error {
	// Connect to the PostgreSQL server
	db, err := utils.DbConnection("")
	if err != nil {
		return err
	}

	// Drop the test database
	dropped := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s WITH (FORCE)", dbName))
	if dropped.Error != nil {
		return dropped.Error
	}
	return nil
}
