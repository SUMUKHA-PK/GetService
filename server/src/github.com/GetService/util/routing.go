package util

import (
	"fmt"
	"net/http"
	"os"

	"../errorAndLog"
)

func RenderPage(w http.ResponseWriter, pageName string) {

	f, err := os.Open(pageName)
	errorAndLog.ErrorHandler(err, "RenderPage, routing.go")
	b1 := make([]byte, 100000)
	_, err = f.Read(b1)
	errorAndLog.ErrorHandler(err, "RenderPage, routing.go")
	fmt.Fprintf(w, string(b1))
}
