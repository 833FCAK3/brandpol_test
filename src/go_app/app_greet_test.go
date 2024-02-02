package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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

	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, Test_db)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	db.AutoMigrate(&Greeting{})
	db.Create(&greeting)

	return c.String(http.StatusOK, greeting.Message)
}
