package main

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	// Set up test database
	testDbName := "postgres_test"
	var err error
	testDB, err = createTestDatabase(testDbName)
	if err != nil {
		fmt.Printf("Error creating test database: %v\n", err)
		os.Exit(1)
	}

	// Run tests
	exitCode := m.Run()

	// // Clean up test database
	// if err := dropTestDatabase(testDB, testDbName); err != nil {
	// 	fmt.Printf("Error dropping test database %v: %v\n", testDbName, err)
	// 	os.Exit(1)
	// }

	os.Exit(exitCode)
}

func createTestDatabase(DbName string) (*gorm.DB, error) {
	// Connect to the PostgreSQL server
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var count int64
	db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", DbName).Scan(&count)
	if count > 0 {
		// Database already exists, return without creating it
		return nil, nil
	}

	// Create the test database
	rawSQL := fmt.Sprintf("CREATE DATABASE %s", DbName)
	if err := db.Exec(rawSQL).Error; err != nil {
		return nil, err
	}

	// Connect to the newly created test database
	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
