package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)



func TestGreetEndpoint(t *testing.T) {
	// Create an Echo instance
	e := echo.New()

	// Create a request to the /greet endpoint (GET request)
	req := httptest.NewRequest(http.MethodGet, "/greet", nil)

	// Create a response recorder to record the response
	rec := httptest.NewRecorder()

	// Handle the request using the Greet endpoint
	c := e.NewContext(req, rec)
	err := handleGreetEndpoint(c)

	// Assert that there was no error handling the request
	assert.NoError(t, err)

	// Assert the HTTP status code is OK (200)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Decode the response body to check the content
	expectedMessage := "Привет от Go!"
	assert.Equal(t, expectedMessage, rec.Body.String())

	// Optionally, you can also check the database for the saved greeting
}

// Function to handle the /greet endpoint
func handleGreetEndpoint(c echo.Context) error {
	// Get the current timestamp
	greeting := Greeting{
		Message:   "Привет от Go!",
		CreatedAt: time.Now(),
	}

	// PostgreSQL connection string
    dsn := "host=localhost user=postgres password=postgres dbname=postgres_test port=5433 sslmode=disable"

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	// Auto-migrate the Greeting table
	db.AutoMigrate(&Greeting{}, &Python_Greet{}, &Python_Greet_History{})

	// Save the greeting to the database
	db.Create(&greeting)

	// Return the greeting message
	return c.String(http.StatusOK, greeting.Message)
}
