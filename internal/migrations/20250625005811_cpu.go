package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCpu, downCpu)
}

func upCpu(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS cpus (
											id INTEGER PRIMARY KEY AUTO_INCREMENT,
											name VARCHAR(255) NOT NULL,
											freq INTEGER NOT NULL,
											temp INTEGER NOT NULL
										)`)
	if err != nil {
		return err
	}
	return nil
}

func downCpu(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS cpus`)
	if err != nil {
		return err
	}
	return nil
}
