package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upToken, downToken)
}

func upToken(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS tokens 
						(id INTEGER AUTO_INCREMENT PRIMARY KEY,
						token VARCHAR(255) NOT NULL,
						FOREIGN KEY (user_id) REFERENCES users(id),`)
	if err != nil {
		return err
	}
	return nil
}

func downToken(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE IF EXISTS tokens;`)
	if err != nil {
		return err
	}
	return nil
}
