package main

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
)

var embedMigrations embed.FS

func migrate() {
	var db *sql.DB

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}

func resetMigrations() {
	var db *sql.DB

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, "migrations"); err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]

	fmt.Println("\033[H\033[2J")
	if len(args) > 0 {
		switch args[0] {
		case "migrate":
			migrate()
		case "migrations":
			migrate()
		case "reset-migrations":
			resetMigrations()
		case "reset-migrate":
			resetMigrations()
		}
	}
}
