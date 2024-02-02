package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetHistoryEndpoint(t *testing.T) {

	response, err := http.Get(fmt.Sprintf("http://%s:8080/greet/history", AppHost))
	if err != nil {
		fmt.Println("Error sending request to Python API")
	}
	defer response.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
