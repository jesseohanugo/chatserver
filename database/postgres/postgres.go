package postgres

import (
	"context"
	"database/sql"

	"github.com/jesseohanugo/chatserver/models"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) *PostgresDB {
	return &PostgresDB{db: db}
}

func (db *PostgresDB) CreateUser(ctx context.Context, user *models.User) error {
	return createUserQuery(ctx, db.db, user) // Use queries.CreateUserQuery function
}

// ... implement other database functionalities based on the interface ...
