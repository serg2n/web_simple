package main

import (
	"log"
	"simple-web-app/constants"
	"simple-web-app/dbstorage/postgres"
	"simple-web-app/http"
)

func main() {
	log.Printf("Simple-Web-Application is running...")

	http.StartServer(constants.SERVER_PORT)
}

func init() {
	postgres.MigrateDbSchema()
}
