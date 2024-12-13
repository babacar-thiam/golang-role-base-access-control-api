package database

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"role-base-access-control-api/internal/role"
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
