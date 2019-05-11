package main

import (
	"fmt"
	"net/http"

	"../routing"
)

func main() {
	go func() {
		mux := &http.ServeMux{}
		mux.HandleFunc("/", routing.HandlerFunc)
		fmt.Print("Web server Serving on port 3000\n")
		http.ListenAndServe("192.168.43.10:3000", mux)
	}()

	fmt.Scanln()
	fmt.Scanln()
	fmt.Print("Server shut down")
}
