package services

import (
	"database/sql"
)

func Check_session_valid(token string, conn *sql.DB) bool {
	query := `SELECT token FROM tokens token = ?`
	if conn.QueryRow(query, token) != nil {
		return false
	}
	return true
}
