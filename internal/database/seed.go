package database

import (
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"role-base-access-control-api/internal/role"
	"role-base-access-control-api/internal/user"
	"time"
)

// SeedRoles creates initial roles if they don't exist
func SeedRoles(db *sql.DB) error {
	// Define default roles
	defaultRoles := []role.Role{
		{
			ID:          uuid.New(),
			Name:        "ADMIN",
			Description: "Administrator access",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "CLIENT",
			Description: "Client access",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "PROVIDER",
			Description: "Provider access",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Initialize role repository
	repo := role.NewRepository(db)

	// Start the transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Check and create each role
	for _, r := range defaultRoles {
		// Check if the role already exists
		existingRole, roleErr := repo.FindByName(r.Name)
		if roleErr != nil {
			log.Printf("Error finding role by name: %v", roleErr)
			return roleErr
		}

		if existingRole != nil {
			log.Printf("Role already exists: %v", existingRole.Name)
			continue
		}

		// Create role if it does not exist
		err = repo.Save(r)
		if err != nil {
			log.Printf("Error saving role: %v", err)
		}

		log.Printf("Created role: %v", r.Name)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	log.Println("role seeding completed successfully")
	return nil
}

// SeedAdmins creates initial admin users if they don't exist
func SeedAdmins(db *sql.DB) error {
	// Define the default admin users
	defaultAdmins := []user.User{
		{
			ID:        uuid.New(),
			Name:      "Admin User",
			Email:     "admin@example.com",
			Phone:     "221771009010",
			Password:  "password",
			Role:      "ADMIN",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Name:      "Super Admin",
			Email:     "super.admin@example.com",
			Phone:     "221771009011",
			Password:  "password",
			Role:      "ADMIN",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Initialize the user repository
	repo := user.NewRepository(db)

	// Start the transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Check and create each admin user
	for _, admin := range defaultAdmins {
		// Check if the admin already exists
		existingAdmin, err := repo.FindByEmail(admin.Email)
		if err != nil {
			log.Printf("Error finding admin by email: %v", err)
			return err
		}

		if existingAdmin != nil {
			log.Printf("Admin already exists: %v", existingAdmin.Email)
			continue
		}

		// Hash the password for the new admin
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if hashErr != nil {
			log.Printf("Error hashing password: %v", err)
			return err
		}
		admin.Password = string(hashedPassword)

		// Save the admin user
		err = repo.Save(&admin)
		if err != nil {
			log.Printf("Error saving admin user: %v", err)
			return err
		}

		log.Printf("Created admin user: %v", admin.Email)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	log.Println("Admin users seeding completed successfully")
	return nil
}
