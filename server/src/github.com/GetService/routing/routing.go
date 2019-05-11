package routing

import (
	"fmt"
	"net/http"

	"../util"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		util.RenderPage(w, "../webpages/dynamic/index.html")
	} else {
		w.WriteHeader(http.StatusNotFound) // Status code 404
		fmt.Fprint(w, "<h1>Error 404 : Page not found</h1>")
	}
}
