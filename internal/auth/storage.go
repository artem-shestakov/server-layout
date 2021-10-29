package auth

import (
	apperror "api/internal/apperrors"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

var (
	rolesTable = "roles"
	userTable  = "users"
)

type Authorization interface {
	CreateUser(user *CreateUser) (*User, error)
	GetUserID(email, password string) (*User, error)
	GetRole(role string) (*Role, error)
}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// CreateUser exec database query
func (s *Storage) CreateUser(user *CreateUser) (*User, error) {
	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, email, password, role) VALUES ($1, $2, $3, $4, $5) RETURNING id, firstname, lastname, email, role, created_at, is_active", userTable)
	row := s.db.QueryRow(query, user.Firstname, user.Lastname, user.Email, user.Password, user.Role)
	var newUser User
	if err := row.Scan(&newUser.Id, &newUser.Firstname, &newUser.Lastname, &newUser.Email, &newUser.Role, &newUser.CreatedAt, &newUser.IsActive); err != nil {
		return nil, apperror.NewError(err, http.StatusInternalServerError, "Create user error", err.Error())
	}
	return &newUser, nil
}

func (s *Storage) GetUserID(email, password string) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", userTable)
	err := s.db.Get(&user, query, email, password)
	if err != nil {
		return nil, apperror.NewError(err, http.StatusNotFound, "User not found", err.Error())
	}
	return &user, nil
}

func (s *Storage) GetRole(roleName string) (*Role, error) {
	var role Role
	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", rolesTable)
	err := s.db.Get(&role, query, roleName)
	if err != nil {
		return nil, apperror.NewError(err, http.StatusNotFound, "Role not found", err.Error())
	}
	return &role, nil
}
