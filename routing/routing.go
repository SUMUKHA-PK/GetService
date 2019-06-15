package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouting adds all the routes
func SetupRouting(r mux.Router) mux.Router {
	r.HandleFunc("/", HomeRouter).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/pasteData", PasteData).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/storePassword", StorePassword).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/getPassword", GetPassword).Methods(http.MethodPost, http.MethodOptions)
	return r
}
