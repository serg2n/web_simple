package http

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(port int) {
	pr := NewPathResolver()
	configureRouting(pr)
	log.Printf("Starting server : localhost:%d", port)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("localhost:%d", port),
			pr))
}
