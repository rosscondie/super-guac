package config

import (
	"errors"
	"log"
	"os"
)

type Credentials struct {
	Username  string
	Password  string
	JWTSecret []byte
}

func LoadCredentials() (Credentials, error) {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	jwtSecret := os.Getenv("JWT_SECRET")

	log.Printf("DEBUG: ADMIN_USERNAME=%q, ADMIN_PASSWORD=%q, JWT_SECRET length=%d\n", username, password, len(jwtSecret))

	if username == "" || password == "" || jwtSecret == "" {
		return Credentials{}, errors.New("missing one or more required environment variables")
	}

	return Credentials{
		Username:  username,
		Password:  password,
		JWTSecret: []byte(jwtSecret),
	}, nil
}
