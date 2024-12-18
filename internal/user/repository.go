package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Repository struct {
	DB *sql.DB
}

// NewRepository handles the database operations
func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// Save saves a new user to the database
func (r *Repository) Save(user *User) error {
	// SQL query to insert a new user to the database
	query := "insert into users (id, name, email, phone, password, role, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)"

	// Execute the query
	_, err := r.DB.Exec(query, user.ID, user.Name, user.Email, user.Phone, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not save user: %w", err)
	}
	return nil
}

// Find finds the user by his ID
func (r *Repository) Find(id uuid.UUID) (*User, error) {
	// Create a variable that hold the user data
	var user User

	// SQL query to retrieve user by email
	query := "select id, name, email, phone, password, role, created_at, updated_at from users where id = ?"

	// Execute the SQL query
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find user by ID: %w", err)
	}

	// Return the found user
	return &user, nil
}

// FindByEmail finds a user by his email
func (r *Repository) FindByEmail(email string) (*User, error) {
	// Create a variable that hold the user data
	var user User

	// SQL query to retrieve user by email
	query := "select id, name, email, phone, password, role, created_at, updated_at from users where email = ?"

	// Execute the SQL query
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find user by email: %w", err)
	}

	// Return the found user
	return &user, nil
}

// FindByPhone finds a user by his phone
func (r *Repository) FindByPhone(phone string) (*User, error) {
	// Create a variable that hold the user data
	var user User

	// SQL query to retrieve user by email
	query := "select id, name, email, phone, password, role, created_at, updated_at from users where phone = ?"

	// Execute the SQL query
	err := r.DB.QueryRow(query, phone).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("could not find user by phone: %w", err)
	}

	// Return the found user
	return &user, nil
}

// FindAll finds all the users
func (r *Repository) FindAll() ([]User, error) {
	// SQL query to fetch all users
	query := "select id, name, email, phone, password, role, created_at, updated_at from users"

	// Execute the query
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not get users: %w", err)
	}

	// Slice to store all users
	var users []User

	for rows.Next() {
		var user User
		scanErr := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if scanErr != nil {
			return nil, fmt.Errorf("could not scan user: %w", scanErr)
		}
		users = append(users, user)
	}

	// Check for errors encounter during iteration
	if rowErr := rows.Err(); rowErr != nil {
		return nil, fmt.Errorf("error getting roles: %v", rowErr)
	}

	return users, nil
}
