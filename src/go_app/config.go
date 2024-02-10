package main

import (
	"os"
)

func getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var GoPort = getenv("GO_PORT", "8080")
var PyPort = getenv("PY_PORT", "8000")
var DbHost = getenv("GO_POSTGRES_HOST", "go_postgres_db")
var DbName = getenv("POSTGRES_DB", "postgres")
var DbUser = getenv("POSTGRES_USER", "postgres")
var DbPassword = getenv("POSTGRES_PASSWORD", "postgres")
var Test_db = getenv("POSTGRES_DB_TEST", "postgres_test")

const (
	AppHost    = "go_app"
	DbPort = 5432
	PyAppHost  = "py_app"
)
