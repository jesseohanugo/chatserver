package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jesseohanugo/chatserver/models"
)

// CreateUserQuery inserts a new user into the database
func createUserQuery(ctx context.Context, db *sql.DB, user *models.User) error {
	query := `
        INSERT INTO auth_schema.users (username, email, password_hash, verified)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	// Replace potential SQL injection vulnerabilities with prepared statement arguments
	row := stmt.QueryRowContext(ctx, user.Username, user.Email, user.PasswordHash, user.Verified)
	err = row.Scan(&user.ID) // Retrieve generated ID if successful
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("duplicate username or email")
		}
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}
