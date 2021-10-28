package auth

import "github.com/jmoiron/sqlx"

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser() error {
	return nil
}
