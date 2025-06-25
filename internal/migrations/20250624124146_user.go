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

	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS users 
						(id INTEGER AUTO_INCREMENT PRIMARY KEY,
						name VARCHAR(255) NOT NULL,
						password VARCHAR(255) NOT NULL,
						role VARCHAR(255) NOT NULL,
						email VARCHAR(255) UNIQUE NOT NULL,
						verified TINYINT(1) DEFAULT 0`) 

	// creating a users table that has an id as a 
	// primary key, a username, a password, encrypted
	// a role (admin, user), a unique email, and a verified status 
	if err != nil {
		return err
	}
	return nil
}

func downUser(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS users;`)

	if err != nil{
		return err
	}
	return nil
}
