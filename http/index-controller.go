package http

import (
	"fmt"
	"net/http"
	"simple-web-app/ui"
)

func  IndexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, ui.IndexHTML)
}

