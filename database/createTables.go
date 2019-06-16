package database

import (
	"database/sql"
	"fmt"

	// To open a PSQL connection
	_ "github.com/lib/pq"
)

// CreateTables returns a DB pointer after creating necessary tables if they dont exist
func CreateTables() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", "localhost", 54320, "postgres", "", "pastebin")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	session := `CREATE TABLE IF NOT EXISTS pastes (
				expiration_length_in_min int NOT NULL, 
				created_at timestamp NOT NULL, 
				pastepath varchar(255) NOT NULL, 
				pastedata text, 
				PRIMARY KEY (pastepath))
				`

	_, err = db.Exec(session)

	return db, err
}
