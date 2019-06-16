package main

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	// pass := "root:" + "newpass" + "@(127.0.0.1:54320)/pastebin"
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable","localhost", 54320, "postgres","", "pastebin")
	// db, err := sql.Open("postgres", psqlInfo)
	// if err!=nil{
	// 	log.Println(err)
	// }

	// session := `CREATE TABLE IF NOT EXISTS pastes (
	// 			hashlink char(255) NOT NULL,
	// 			expiration_length_in_min int NOT NULL,
	// 			created_at timestamp NOT NULL,
	// 			pastepath varchar(255) NOT NULL,
	// 			PRIMARY KEY (hashlink))
	// 			`

	// _, err = db.Exec(session)
	// if err!=nil{
	// 	log.Println(err)
	// }

	// defer db.Close()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
