package auth

import (
	"database/sql"
	"errors"

	"synf/internal/api/data/models"
	"synf/internal/database"

	"github.com/google/uuid"
)

func CreateSession(role string) (string, error) {
	token := uuid.NewString()
	query := `INSERT INTO sessions (token) VALUE ? `

	conn, err := database.Connect()
	if err != nil {
		return token, err
	}

	_, err = conn.Exec(query, token)
	if err != nil {
		return token, err
	}

	return token, nil
}

func ValidateCredentials(user models.User) (bool, error) {
	conn, err := database.Connect()
	if err != nil {
		return false, err
	}

	defer database.Close(conn)
	query := ` SELECT * FROM users WHERE users.id = ?`

	// QueryRow expects at most one row
	var id uint
	err = conn.QueryRow(query, user.UserID).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
