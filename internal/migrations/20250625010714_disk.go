package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upDisk, downDisk)
}

func upDisk(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS disks (
											id INTEGER PRIMARY KEY AUTO_INCREMENT,
											name VARCHAR(255) NOT NULL,
											size INTEGER NOT NULL
										)`)
	if err != nil {
		return err
	}

	return nil
}

func downDisk(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE IF EXISTS disks`)
	if err != nil {
		return err
	}
	return nil
}
