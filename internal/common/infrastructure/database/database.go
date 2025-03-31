package database

import (
	"database/sql"
	"fmt"
	"os"

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
		panic(err)
	} else {
		fmt.Println("Database connected")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	createTables(db)

	return &DB{
		Client: db,
	}, nil
}

func createTables(db *sql.DB) {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			password TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`

	_, err := db.Exec(createUsersTable)

	if err != nil {
		fmt.Println("Create tables failed")
		panic(err)
	}
}
