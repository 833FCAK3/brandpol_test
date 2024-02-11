package main

import (
	"log"
	"myapp/config"
	"myapp/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGreetEndpoint(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := handleGreetEndpoint(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedMessage := "Привет от Go!"
	assert.Equal(t, expectedMessage, rec.Body.String())
}

func handleGreetEndpoint(c echo.Context) error {
	// Get the current timestamp
	greeting := Greeting{
		Message:   "Привет от Go!",
		CreatedAt: time.Now(),
	}

	// Connect to the test database
	db, err := utils.DbConnection(config.Test_db)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&Greeting{})
	db.Create(&greeting)

	return c.String(http.StatusOK, greeting.Message)
}
