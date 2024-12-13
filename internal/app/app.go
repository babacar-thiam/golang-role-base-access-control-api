package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"role-base-access-control-api/configs"
	"role-base-access-control-api/internal/database"
	"role-base-access-control-api/internal/role"
)

type App struct {
	Config *configs.Config
	DB     *sql.DB
	Router *mux.Router
}

// NewApp creates a new App instance
func NewApp(cfg *configs.Config) (*App, error) {
	// Initialize the database connection
	conn, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}

	// Seed the default roles
	if seedErr := database.SeedRoles(conn); seedErr != nil {
		log.Fatalf("failed to seed roles: %v", seedErr)
	}

	// Create App instance
	app := &App{
		Config: cfg,
		DB:     conn,
		Router: mux.NewRouter(),
	}

	// Initialize routes
	app.setupRoutes()

	return app, nil
}

// setupRoutes defines all application routes
func (app *App) setupRoutes() {
	// Add an API prefix
	api := app.Router.PathPrefix("/api/v1").Subrouter()

	// Initialize role handler
	roleRepo := role.NewRepository(app.DB)
	roleService := role.NewService(roleRepo)
	roleHandler := role.NewHandler(roleService)

	// Define role routes
	api.HandleFunc("/roles", roleHandler.GetRoles).Methods(http.MethodGet)
}

// RunApp starts the HTTP server
func (app *App) RunApp() {
	defer func(DB *sql.DB) {
		_ = DB.Close()
	}(app.DB)

	port := app.Config.AppPort
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}
