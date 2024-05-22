package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/jesseohanugo/chatserver/database/postgres"
	"github.com/jesseohanugo/chatserver/models"
)

// RegisterHandler handles user registration requests
func RegisterHandler(c *gin.Context) {
	var registrationDetails models.RegistrationDetails

	// Use BindJSON to bind request body to the registration model
	if err := c.BindJSON(&registrationDetails); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate user data (optional)
	// ... validation logic ...

	// Hash the password
	hashedPassword, err := hashPassword(registrationDetails.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create a new User object for database
	user := models.User{
		Username:     registrationDetails.Username,
		Email:        registrationDetails.Email,
		PasswordHash: hashedPassword,
		Verified:     false, // Set initial verification state
		// ... set other default values ...
	}

	// ... (Optional) Generate email verification token and set expiry ...

	// Access the database handle from gin.Context
	dbInterface, exists := c.Get("postgresDB")

	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Database handle not found"})
		return
	}

	// Perform type assertion to check if the database handle is of type *postgres.PostgresDB
	postgresDB, ok := dbInterface.(*postgres.PostgresDB)

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Database handle is of the wrong type"})
		return
	}

	ctx := context.Background()

	err = postgresDB.CreateUser(ctx, &user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Handle successful registration (optional)
	// ... send verification email, redirect, etc. ...

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// hashPassword uses bcrypt to hash a password
func hashPassword(password string) ([]byte, error) {
	// Generate a salt with desired cost (14 is recommended for bcrypt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error generating hash: %w", err)
	}

	return hashedPassword, nil
}
