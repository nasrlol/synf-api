package main

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

var embedMigrations embed.FS

func migrations() {
	var db *sql.DB

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("\033[H\033[2J")
}
