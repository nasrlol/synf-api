package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upDevice, downDevice)
}

func upDevice(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS devices (
											id INTEGER PRIMARY KEY AUTO_INCREMENT,
											name VARCHAR(255) NOT NULL,
											status TINYINT(1) DEFAULT 0,
											is_logging TINYINT(1) DEFAULT 0,
											user_id INTEGER,
											cpu_id INTEGER,
											ram_id INTEGER,
											disk_id INTEGER,
											FOREIGN KEY (user_id) REFERENCES users(id),
											FOREIGN KEY (cpu_id) REFERENCES cpus(id),
											FOREIGN KEY (ram_id) REFERENCES rams(id),
											FOREIGN KEY (disk_id) REFERENCES disks(id)
											)`)
	if err != nil {
		return err
	}
	return nil
}

func downDevice(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS devices`)
	if err != nil {
		return err
	}
	return nil
}
