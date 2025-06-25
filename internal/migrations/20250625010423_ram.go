package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upRam, downRam)
}

func upRam(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE rams (
											id INTEGER PRIMARY KEY AUTO_INCREMENT,
											size INTEGER NOT NULL,
											speed INTEGER
											)`)

	if err != nil {
		return err
	}
	return nil
}

func downRam(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
