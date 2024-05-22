package models

import (
	"time"
)

// User represents a user in the application
type User struct {
	ID               int        `json:"id"`
	Username         string     `json:"username"`
	Email            string     `json:"email"`
	PasswordHash     []byte     `json:"password_hash,omitempty"` // Exclude password from JSON serialization
	Verified         bool       `json:"verified"`
	ResetToken       string     `json:"reset_token,omitempty"`
	ResetTokenExpiry *time.Time `json:"reset_token_expiry,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
