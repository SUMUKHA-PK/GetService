package database

import (
	"log"
	"time"

	"github.com/SUMUKHA-PK/GetService/util"
)

// StorePasteData adds the paste data into the DB
func StorePasteData(hashedString string, data util.PasteRequest) error {

	dbConn, err := CreateTables()
	if err != nil {
		return err
	}

	query := "INSERT INTO pastes VALUES ('" +
		data.ExpirationTime + "','" +
		time.Now().Format("2006-01-02 15:04:05") + "','" +
		hashedString + "','" +
		data.PasteContent + "'" +
		")"

	log.Println(query)
	_, err = dbConn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
