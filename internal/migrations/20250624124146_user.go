package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUser, downUser)
}

func upUser(ctx context.Context, tx *sql.Tx) error {

	_, err := tx.Exec(`CREATE TABLE users (
											id INTEGER PRIMARY KEY AUTO_INCREMENT,
											name VARCHAR(255) NOT NULL,
											email VARCHAR(255) NOT NULL UNIQUE,
											password VARCHAR(255) NOT NULL,
											role VARCHAR(255) NOT NULL,
											verified BOOLEAN NOT NULL
										)`)

	if err != nil {
		return err
	}
	return nil
}

func downUser(ctx context.Context, tx *sql.Tx) error {

	_, err := tx.Exec(`DROP TABLE IF EXISTS users`)
	if err != nil {
		return err
	}
	return nil
}
