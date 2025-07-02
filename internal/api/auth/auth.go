package auth

import (
	"fmt"

	reg "synf/internal/api/data/rest"
	db "synf/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

func GetUser(userName string) (reg.UserInformation, error) {
	var userInfo reg.UserInformation
	query := "SELECT user.id FROM user WHERE name LIKE ?"

	conn, err := db.ConnectDB()
	if err != nil {
		return reg.UserInformation{}, fmt.Errorf("failed to connect to the database")
	}

	row := conn.QueryRow(query, userName)
	err = row.Scan(&userInfo.UserName)
	if err != nil {
		return reg.UserInformation{}, fmt.Errorf("there was an error in retrieving the information from the database")
	}
	return userInfo, nil
}

func CheckPass(pass string, data reg.UserInformation) bool {
	query := "SELECT ? FROM ? WHERE ? LIKE ?"

	conn, err := db.ConnectDB()
	if err != nil {
		return false
	}

	queryError := conn.QueryRow(query, data.UserID, "USER", data.UserPassword, pass)
	return queryError == nil
}
