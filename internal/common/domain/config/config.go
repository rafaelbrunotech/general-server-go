package config

import (
	"os"
)

var Env = map[string]string{
	"ENV":         os.Getenv("ENV"),
	"JWT_SECRET":  os.Getenv("JWT_SECRET"),
	"PORT":        os.Getenv("PORT"),
	"DB_DRIVER":   os.Getenv("DB_DRIVER"),
	"DB_HOST":     os.Getenv("DB_HOST"),
	"DB_NAME":     os.Getenv("DB_NAME"),
	"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
	"DB_PORT":     os.Getenv("DB_PORT"),
	"DB_USER":     os.Getenv("DB_USER"),
}
