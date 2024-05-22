package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jesseohanugo/chatserver/config" // config package
	_ "github.com/lib/pq"                       // PostgreSQL driver
)

// Connect establishes a connection to the PostgreSQL database
func Connect() (*sql.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL database successfully!")
	return db, nil
}
