package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/pressly/goose/v3"

	h "synf/internal/api/routes"
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

func RawConnect(host string, port string) {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Connecting error:", err)
	}
	if conn != nil {
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(conn)
		fmt.Println("Opened", net.JoinHostPort(host, port))
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
		case "start":

			fmt.Println("\033[H\033[2J")
			r := h.InitRestRoutes()
			w := h.InitWsRoutes()
			log.Fatal(http.ListenAndServe(":8080", r))
			log.Fatal(http.ListenAndServe(":8090", w))
		}
	}
}
