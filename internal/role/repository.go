package role

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Repository struct {
	DB *sql.DB
}

// NewRepository handles roles database operations
func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// Save inserts a new role to the database
func (r *Repository) Save(role Role) error {
	// SQL query to insert a new role
	query := "insert into roles (id, name, description, created_at, updated_at) values (?, ?, ?, ?, ?)"

	// Execute the query
	_, err := r.DB.Exec(query, role.ID, role.Name, role.Description, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("error saving role: %v", err)
	}
	return nil
}

func (r *Repository) FindByName(name string) (*Role, error) {
	// SQL query to retrieve a role by its name
	query := "select id, name, description, created_at, updated_at from roles where name = ?"

	// Create a variable to hold the role data
	var role Role

	// Execute the query
	err := r.DB.QueryRow(query, name).Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Return nil for both role and error if no role is found
			return nil, nil
		}
		return nil, fmt.Errorf("error getting role: %v", err)
	}

	// Return the found role
	return &role, nil
}

// FindAll retrieve all the roles from the database
func (r *Repository) FindAll() ([]Role, error) {
	// SQL query to retrieve all the roles
	query := "select id, name, description, created_at, updated_at from roles"

	// Execute the query
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting roles: %v", err)
	}

	// Create a variable that holds all roles data
	var roles []Role

	for rows.Next() {
		var role Role
		// Scan each role into a Role struct
		scanErr := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if scanErr != nil {
			return nil, fmt.Errorf("error getting roles: %v", scanErr)
		}
		roles = append(roles, role)
	}

	// Check for errors encounter during iteration
	if rowErr := rows.Err(); rowErr != nil {
		return nil, fmt.Errorf("error getting roles: %v", rowErr)
	}

	// Return the found roles
	return roles, nil
}
