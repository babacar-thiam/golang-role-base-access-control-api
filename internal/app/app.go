package app

import (
	"database/sql"
	"log"
	"net/http"
	"role-base-access-control-api/configs"
	"role-base-access-control-api/internal/database"
)

type App struct {
	Config *configs.Config
	DB     *sql.DB
}

// NewApp creates a new App instance
func NewApp(cfg *configs.Config) (*App, error) {
	// Initialize the database connection
	conn, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		DB:     conn,
	}, nil
}

// RunApp starts the application
func (app *App) RunApp() error {
	defer func(DB *sql.DB) {
		_ = DB.Close()
	}(app.DB) // Ensure the database is closed when app shuts down

	log.Printf("Starting server on port %s...", app.Config.AppPort)
	return http.ListenAndServe(":"+app.Config.AppPort, nil)
}
