package http

import (
	"fmt"
	"log"
	"net/http"
	"simple-web-app/dbstorage/postgres"
)

func StartServer(port int) {

	//Dependency Injection
	cs := postgres.NewContactServiceImpl(postgres.DbConnection())
	cc := NewContactController(cs)
	router := NewRouter(cc)
	pr := NewPathResolver()
	router.configureRouting(pr)

	log.Printf("Starting server : 0.0.0.0:%d", port)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("0.0.0.0:%d", port),
			pr))
}
