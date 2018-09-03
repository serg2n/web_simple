package controller

import "net/http"

func Contact(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("content-type", "application/json")
	res.Write([]byte("{ \"result\":\"OK\"}"))
}
