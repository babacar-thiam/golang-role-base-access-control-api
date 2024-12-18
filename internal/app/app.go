package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"role-base-access-control-api/api/v1"
	"role-base-access-control-api/configs"
	_ "role-base-access-control-api/docs"
	"role-base-access-control-api/internal/auth"
	"role-base-access-control-api/internal/database"
	"role-base-access-control-api/internal/role"
	"role-base-access-control-api/internal/user"
)

type App struct {
	Config *configs.Config
	DB     *sql.DB
	Router *mux.Router
	api    *v1.API
}

func NewApp(cfg *configs.Config) (*App, error) {
	// Initialize database connection
	conn, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}

	// Seed database
	if err := database.SeedRoles(conn); err != nil {
		log.Fatalf("failed to seed roles: %v", err)
	}
	if err := database.SeedAdmins(conn); err != nil {
		log.Fatalf("failed to seed admins: %v", err)
	}

	// Initialize repositories
	roleRepo := role.NewRepository(conn)
	userRepo := user.NewRepository(conn)

	// Initialize services
	roleService := role.NewService(roleRepo)
	userService := user.NewService(userRepo)
	authService := auth.NewService(cfg, userRepo)

	// Initialize JWT and middleware
	jwt := auth.NewJWT(cfg)
	authMiddleware := auth.NewMiddleware(jwt)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)
	roleHandler := role.NewHandler(roleService)
	userHandler := user.NewHandler(userService)

	// Create router
	router := mux.NewRouter()

	// Initialize API
	api := v1.NewAPI(authHandler, roleHandler, userHandler, authMiddleware)

	// Create App instance
	app := &App{
		Config: cfg,
		DB:     conn,
		Router: router,
		api:    api,
	}

	// Setup routes
	app.setupRoutes()

	// Swagger UI route
	app.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return app, nil
}

func (app *App) setupRoutes() {
	// Setup API routes
	app.api.SetupRoutes(app.Router)
}

func (app *App) RunApp() {
	port := app.Config.AppPort
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}
