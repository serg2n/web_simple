package main

import (
	"log"
	"simple-web-app/dbstorage"
)

func main() {
	log.Printf("Simple-Web-Application is running...")
	log.Printf("Simple-Web-Application is shutting down...")
}

func init() {
	dbstorage.MigrateDbSchema()
}
