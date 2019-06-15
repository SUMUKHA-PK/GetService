package routing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	random "math/rand"

	"github.com/SUMUKHA-PK/Database-password-management-system/crypto"
	"github.com/SUMUKHA-PK/Database-password-management-system/util"
)

// PasteData handles /pasteData
func PasteData(w http.ResponseWriter, r *http.Request) {

	// get the data from the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtain the JSON data in a struct
	var newReq PasteRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//check for duplicates
	clientIP := r.RemoteAddr
	dateTimeStamp := time.Now().String()
	preHashedString := clientIP + dateTimeStamp + util.StringWithCharset(random.Intn(20)+10, util.Charset)
	hashedString := crypto.Create

	URL := PasteResponse{"ITH"}
	outJSON, err := json.Marshal(URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(outJSON)
}
