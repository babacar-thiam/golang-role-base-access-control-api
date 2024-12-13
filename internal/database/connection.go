package database

import (
	"database/sql"
	"fmt"
	"log"
	"role-base-access-control-api/configs"

	_ "github.com/go-sql-driver/mysql"
)

// Connect creates and return a new MySQL connecting u
// using the provided config
func Connect(cfg *configs.Config) (*sql.DB, error) {
	// Define the database source name(DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Open the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")
	return db, nil
}
