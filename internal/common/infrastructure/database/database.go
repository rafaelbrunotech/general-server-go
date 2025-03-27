package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rafaelbrunoss/general-server-go/internal/common/domain/config"
)

type DB struct {
	Client *sql.DB
}

func InitDB() (*DB, error) {
	driver := config.Env["DB_DRIVER"]
	user := config.Env["DB_USER"]
	name := config.Env["DB_NAME"]
	password := config.Env["DB_PASSWORD"]
	host := config.Env["DB_HOST"]
	port := config.Env["DB_PORT"]

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
			name VARCHAR(255) NOT NULL,
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
