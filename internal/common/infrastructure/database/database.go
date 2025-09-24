package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	Client *sql.DB
}

func InitDB() (*DB, error) {
	driver := os.Getenv("DB_DRIVER")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)

	db, err := sql.Open(driver, connection)

	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	fmt.Println("Database connected")

	return &DB{
		Client: db,
	}, nil
}

func (db *DB) Close() error {
	return db.Client.Close()
}
