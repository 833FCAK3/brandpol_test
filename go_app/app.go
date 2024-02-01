package main

import (
	"encoding/json"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/labstack/echo/v4"
	"fmt"
	"log"
    "os"
    "strconv"
    "net/http"
    "time"
	"io"
)

type Greeting struct {
	// Exclude the default fields by using struct tags
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Message   string    `json:"message"`
}

type Python_Greet struct {
	gorm.Model
	Name      string `json:"name"`
	PythonGreeting string `json:"python_greet"`
}

type Python_Greet_History struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PythonGreetings string 	`json:"greetings"`
	CreatedAt time.Time
}

func main() {
    // Retrieve port from the environment variables
    portStr := os.Getenv("GO_PORT")
    if portStr == "" {
        portStr = "8080" // Default port if not specified in .env
    }

    // Convert the port string to an integer
    port, err := strconv.Atoi(portStr)
    if err != nil {
        log.Fatal("Invalid PORT value in .env")
    }

    // PostgreSQL connection string
    dsn := "host=go_postgres_db user=postgres password=postgres dbname=postgres port=5433 sslmode=disable port=5432"

    // Connect to the database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database" + err.Error())
    }

    // Auto-migrate the Greeting table
    db.AutoMigrate(&Greeting{}, &Python_Greet{}, &Python_Greet_History{})

    // Create an Echo instance
    e := echo.New()

    // Endpoint to greet and save timestamp to the database
    e.GET("/greet", func(c echo.Context) error {
        // Save timestamp to the database
        greeting := Greeting{
            Message:   "Привет от Go!",
            CreatedAt: time.Now(),
        }
        db.Create(&greeting)

        // Return the greeting message
        return c.String(http.StatusOK, greeting.Message)
    })

	// Endpoint to retrieve all greetings from the database
	e.GET("/greet/history", func(c echo.Context) error {
		// Query all greeting entries
		var greetings []Greeting
		db.Order("created_at desc").Find(&greetings)

		// Return the greeting history
		return c.JSON(http.StatusOK, greetings)
	})

	// Endpoint to reach python greet and save response to db
	e.GET("/python_greet", func(c echo.Context) error {
		// Get the 'name' parameter from the request
		name := c.QueryParam("name")
		if name == "" {
			return c.String(http.StatusBadRequest, "Name parameter is required")
		}

		// Specify the URL to send the GET request to
		pythonAPIURL := "http://localhost:8000/greet"

		// Send GET request to Python API
		response, err := http.Get(fmt.Sprintf("%s?name=%s", pythonAPIURL, name))
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error sending request to Python API")
		}
		defer response.Body.Close()

		// Decode the JSON response from Python API
		var pythonGreetResponse map[string]string
		err = json.NewDecoder(response.Body).Decode(&pythonGreetResponse)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error decoding Python API response")
		}

		// Save the received Python greeting to the database
		pythonGreeting := pythonGreetResponse["message"]
		greeting := Python_Greet{
			Name:      name,
			PythonGreeting: pythonGreeting,
		}
		db.Create(&greeting)

		// Return the Python greeting
		return c.String(http.StatusOK, pythonGreeting)
	})

	// Endpoint to reach python greet history and save response to db
	e.GET("/python_greet_history", func(c echo.Context) error {

		// Specify the URL to send the GET request to
		pythonAPIURL := "http://localhost:8000/greet/history"

		// Send GET request to Python API
		response, err := http.Get(pythonAPIURL)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error sending request to Python API")
		}
		defer response.Body.Close()
		// if err != nil {
		// 	errorMsg := fmt.Sprintf("Error sending request to Python API: %s", err.Error())
		// 	return c.String(http.StatusInternalServerError, errorMsg)
		// }
		// defer response.Body.Close()

		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		// Save the entire response body as a single entry in the Echo application's database
		data := string(body)
		greetHistory := Python_Greet_History{
			PythonGreetings: data,
		}
		db.Create(&greetHistory)

		// Return a response
		return c.String(http.StatusOK, data)
	})

    e.Start(":" + strconv.Itoa(port))
}
