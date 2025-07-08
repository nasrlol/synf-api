package database

import (
	"database/sql"
	"fmt"
	"time"

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

	// Check connection
	if err = db.Ping(); err != nil {
		fmt.Println("failed to ping the server")
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Connected to MySQL")
	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}
