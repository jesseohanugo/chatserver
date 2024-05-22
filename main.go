package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// local imports
	"github.com/jesseohanugo/chatserver/database/postgres"
	"github.com/jesseohanugo/chatserver/handlers/auth"
	"github.com/jesseohanugo/chatserver/middleware"
)

func main() {

	// Initialize connection to database
	db, err := postgres.Connect()

	if err != nil {
		log.Fatal("Error connecting to the database server", err)
	}

	// Create a concrete implementation of the Database interface
	postgresDB := postgres.NewPostgresDB(db)

	// Create a Gin router
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Define a route for the root path ("/")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Define a route for the registration endpoint
	router.POST("/api/v1/register", middleware.DatabaseMiddleware(postgresDB), auth.RegisterHandler)

	// Define a route for the login endpoint
	router.POST("/api/v1/login", auth.LoginHandler)

	// Define a route for the email verification endpoint
	router.POST("api/v1/email/verify", middleware.DatabaseMiddleware(postgresDB), auth.RegisterHandler)

	// Start the server on port 8080
	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
