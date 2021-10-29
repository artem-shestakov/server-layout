package config

import "time"

type Config struct {
	Database      map[string]string `yaml:"database"`
	Authorization Authorization
}

type Authorization struct {
	PasswordSalt string
	JWTKey       string
	JWTTTL       time.Duration
}
