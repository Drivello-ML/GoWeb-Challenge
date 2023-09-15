package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GoWeb-Challenge/cmd/server/routes"
	utils "github.com/GoWeb-Challenge/package"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Logger middleware
func LoggerMiddleware(c *gin.Context) {
	// Log the request details
	println("Request URL:", c.Request.URL.String())
	// Continue to the next middleware or handler
	c.Next()
}

// AuthMiddleware is the authentication middleware
func AuthMiddleware(c *gin.Context) {
	authString := c.GetHeader("Authorization")
	token := os.Getenv("TOKEN")
	// Check if the API key exists in the validAPIKeys map
	if authString == token {
		// Valid API key, continue with the next middleware or handler
		c.Next()
	} else {
		// Invalid API key, return a 401 Unauthorized response
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort() // Stop further processing
	}
}

func main() {

	// Load csv.
	loadedTasks, errCsv := utils.LoadTasksCsv("../../tasks.csv")
	if errCsv != nil {
		log.Fatal("Couldn't load tasks.")
		panic(errCsv)
	} else {
		log.Print("tasks loaded successfully")
	}

	errDotEnv := godotenv.Load("../../.env")
	if errDotEnv != nil {
		log.Fatal("Error loading environment variables")
		panic(errDotEnv)
	}

	user := os.Getenv("USER")

	server := gin.Default()

	server.Use(LoggerMiddleware)
	server.Use(AuthMiddleware)

	router := routes.NewTaskRouter(server, loadedTasks)
	router.LoadRoutes()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Greetings %s!", user),
		})
	})

	server.Run(":8080")

}
