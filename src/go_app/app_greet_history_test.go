package main

import (
	"fmt"
	"log"
	"myapp/config"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetHistoryEndpoint(t *testing.T) {

	response, err := http.Get(fmt.Sprintf("http://%s:%s/greet/history", config.PyAppHost, config.PyPort))
	if err != nil {
		log.Fatalf("Error sending request to Python API: %v", err)
	}
	defer response.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
