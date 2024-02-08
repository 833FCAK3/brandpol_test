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
	var err error
	testDB, err = createTestDatabase(Test_db)
	if err != nil {
		fmt.Printf("Error creating test database: %v\n", err)
		os.Exit(1)
	}

	// Run tests
	exitCode := m.Run()

	// Clean up test database
	if err := dropTestDatabase(Test_db); err != nil {
		fmt.Printf("Error dropping test database %v: %v\n", Test_db, err)
		os.Exit(1)
	}

	os.Exit(exitCode)
}

func createTestDatabase(dbname string) (*gorm.DB, error) {
	// Connect to the PostgreSQL server
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func dropTestDatabase(dbName string) error {
	// Connect to the PostgreSQL server
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Drop the test database
	result := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s WITH (FORCE)", dbName))
	if result.Error != nil {
		fmt.Printf("Error dropping test database %v: \n", dbName)
		return result.Error
	}
	return nil
}
