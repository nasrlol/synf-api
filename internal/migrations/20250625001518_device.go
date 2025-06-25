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
	_, err := tx.Exec(`CREATE TABLE devices (
											id INTEGER PRIMARY KEY AUTO_INCREMENT,
											name VARCHAR(255) NOT NULL,
											status TINYINT(1) DEFAULT 0,
											isLogging TINYINT(1) DEFAULT 0
											USERid INTEGER,
											CPUid INTEGER,
											RAMid INTEGER,
											DISKid INTEGER,
											FOREIGN KEY (USERid) REFERENCES users(id)
											FOREIGN KEY (CPUid) REFERENCES cpus(id)
											FOREIGN KEY (RAMid) REFERENCES rams(id)
											FOREIGN KEY (DISKid) REFERENCES disks(id)
											)`)
	if err != nil {
		return err
	}
	return nil
}

func downDevice(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE device IF EXISTS`)
	if err != nil {
		return err
	}
	return nil
}
