package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Logger   logrus.Logger
}

func New(db *Database) (*sqlx.DB, error) {
	database, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		db.Host, db.Port, db.User, db.DBName, db.Password, db.SSLMode))
	if err != nil {
		db.Logger.Fatalf("Can't connect to database")
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		db.Logger.Fatalf("Can't ping database")
		return nil, err
	}
	db.Logger.Infof("Connection to database %s is successul", db.Host)
	return database, nil
}
