package auth

import (
	reg "synf/internal/api/registration"
	db "synf/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

func GetUser(userName string) reg.UserInformation {

	var userInfo reg.UserInformation
	query := "SELECT user.id FROM user WHERE name LIKE ?"

	conn, err := db.ConnectDB()
	if err != nil {
		return reg.UserInformation{}
	}

	row := conn.QueryRow(query, userName)
	err = row.Scan(&userInfo.UserName)

	return userInfo
}

func CheckPass(pass string, data reg.UserInformation) bool {

	query := "SELECT ? FROM ? WHERE ? LIKE ?"

	conn, err := db.ConnectDB()
	if err != nil {
		return false
	}

	queryError := conn.QueryRow(query, data.UserID, "USER", data.UserPassword, pass)
	if queryError != nil {
		return false
	}
	return true
}
