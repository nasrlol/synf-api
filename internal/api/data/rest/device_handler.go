package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	ser "synf/internal/api/data/services"
	db "synf/internal/database"
)

func GetDevice(w http.ResponseWriter, r *http.Request, i ser.DiskInformation) error {
	conn, err := db.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to the database")
	}

	query := `SELECT * FROM disks WHERE disk.id = ?`
	data, err := conn.Exec(query, i.DiskID)
	if err != nil {
		fmt.Errorf("failed to retrieve the data from the database")
	}

	w.Write(json.NewDecoder(data))

}
