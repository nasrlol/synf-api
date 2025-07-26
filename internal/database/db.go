package database

import (
	"database/sql"
	"fmt"

	"synf/internal/config"
)

func Connect() (*sql.DB, error) {
	credentials := config.LoadCredetials()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", credentials.User, credentials.Pass, credentials.Ip, credentials.Port, credentials.Name)

	fmt.Println(credentials.Ip, credentials.User, credentials.Ip, credentials.Port, credentials.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database")
	}

	fmt.Println("Connected to database")
	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}
