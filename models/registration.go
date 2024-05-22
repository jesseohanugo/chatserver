package models

// RegistrationDetails represents user data for registration
type RegistrationDetails struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
