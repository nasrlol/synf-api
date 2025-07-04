package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upSession, downSession)
}

func upSession(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS devices (
											id INTEGER PRIMARY KEY AUTO_INCREMENT,
											token VARCHAR(255) NOT NULL
											)`)
	if err != nil {
		return err
	}
	return nil
}

func downSession(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS sessions`)
	if err != nil {
		return err
	}
	return nil
}
