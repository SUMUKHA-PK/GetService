package database

import (
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

	_, err = dbConn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// GetPasteData gets the pasted data from the DB
func GetPasteData(url string) (string, error) {

	dbConn, err := CreateTables()
	if err != nil {
		return "", err
	}

	data := make([]util.PasteData, 0)
	query := "SELECT * FROM pastes WHERE pastepath = '" + url + "'"

	rows, err := dbConn.Query(query)
	defer rows.Close()
	for rows.Next() {
		var tempData util.PasteData
		if err = rows.Scan(&tempData.ExpLength, &tempData.CreatedTime, &tempData.PastePath, &tempData.Data); err != nil {
			return "", err
		}
		data = append(data, tempData)
	}
	if err := rows.Err(); err != nil {
		return "", err
	}

	return data[0].Data, nil
}
