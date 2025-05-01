package auth

import (
	db "synf/database"
	_ "github.com/go-sql-driver/mysql"
	reg "synf/api/registration"
)

func GetUser(userName string) reg.UserInformation{

	var	userInfo reg.UserInformation
	conn, err := db.ConnectDB()	
	if err != nil {
		return reg.UserInformation{}
	}
	query := "SELECT * FROM user WHERE name LIKE ?"
	row := conn.QueryRow(query, string(userName))
	err = row.Scan(&userInfo.UserName)

	return userInfo
}

func userNameInput() string {
	
	return "hello world"
}

func checkPass(encryptedPass string) bool{

	return true
}


