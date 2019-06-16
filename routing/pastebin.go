package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unsafe"

	random "math/rand"

	"github.com/SUMUKHA-PK/Database-password-management-system/crypto"
	"github.com/SUMUKHA-PK/Database-password-management-system/util"
	"github.com/SUMUKHA-PK/GetService/database"
	Util "github.com/SUMUKHA-PK/GetService/util"
)

// PasteData handles /pasteData
func PasteData(w http.ResponseWriter, r *http.Request) {

	// get the data from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/pastebin.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtain the JSON data in a struct
	var newReq Util.PasteRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Coudn't Unmarshal data in routing/pastebin.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// max size of the paste can be 1KB
	if unsafe.Sizeof(newReq.PasteContent) > 1000 {
		http.Error(w, "Paste size too large!", http.StatusInternalServerError)
		return
	}

	// provide an option for custom strings
	var hashedString string
	duplicate, err := CheckDuplicateURL(newReq.CustomURL)
	if err != nil {
		log.Printf("DB error in routing/pastebin.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if newReq.CustomURL != "" && !duplicate {
		hashedString = Util.SanitizeURL(newReq.CustomURL)
	} else {
		// use client IP, the current time and a random string to salt the hash
		// MD5 hashing is done as used in imported URL
		clientIP := r.RemoteAddr
		dateTimeStamp := time.Now().String()
		preHashedString := clientIP + dateTimeStamp + util.StringWithCharset(random.Intn(20)+10, util.Charset)
		hashedString = crypto.CreateMD5Hash(preHashedString)
	}

	// store the data in the database
	err = database.StorePasteData(hashedString, newReq)
	if err != nil {
		log.Printf("Can't add to DB in routing/pastebin.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// the data is sent back to the WebServer as a JSON
	URL := Util.PasteResponse{hashedString}
	outJSON, err := json.Marshal(URL)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/pastebin.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(outJSON)
}

// func ReadPaste(w http.ResponseWriter, r *http.Request) {

// }

// CheckDuplicateURL returns whether the current URL is taken or not
func CheckDuplicateURL(url string) (bool, error) {

	dbConn, err := database.CreateTables()
	if err != nil {
		return false, err
	}

	query := "SELECT * FROM pastes WHERE pastepath = '" + url + "'"

	rows, err := dbConn.Query(query)
	if err != nil {
		return false, err
	}

	data := make([]string, 0)

	for rows.Next() {
		var expiration_length_in_min, created_at, pastepath, pastedata string
		if err := rows.Scan(&expiration_length_in_min, &created_at, &pastepath, &pastedata); err != nil {
			log.Fatal(err)
		}
		data = append(data, expiration_length_in_min)
	}

	if err := rows.Err(); err != nil {
		return false, err
	}

	if err != nil || len(data) == 0 {
		return false, err
	}

	return true, nil
}
