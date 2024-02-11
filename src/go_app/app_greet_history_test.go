package main

import (
	"fmt"
	"net/http"
	"myapp/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetHistoryEndpoint(t *testing.T) {

	response, err := http.Get(fmt.Sprintf("http://%s:%s/greet/history", config.PyAppHost, config.PyPort))
	if err != nil {
		fmt.Println("Error sending request to Python API")
	}
	defer response.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
