package dbstorage

import (
	"database/sql"
	"fmt"
	"github.com/DavidHuie/gomigrate"
	"log"
	"simple-web-app/constants"
	_ "github.com/lib/pq"
)

var dbConnection *sql.DB = nil

func DbConnection() (*sql.DB) {
	if dbConnection == nil {
		dbConnection = createDbConnection()
	}
	return dbConnection
}

func createDbConnection() (*sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		constants.HOST, constants.PORT, constants.USER, constants.PASSWORD, constants.DbName)

	log.Printf("Connecting to database: %s", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Can not connect to database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Can not ping database: %v", err)
	}

	log.Println("Successfully connected!")
	return db
}

func NextSeqVal() (int32, error) {
	var seqVal int32
	err := DbConnection().QueryRow("SELECT nextval('seq')").Scan(&seqVal)
	if err != nil {
		return -1, nil
	}
	return seqVal, nil
}

func MigrateDbSchema() {
	db := DbConnection()
	migrator, err := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./dbstorage/migrations/")
	if err != nil {
		log.Fatalf("Can not migrate db schema: %v", err)
	}
	err = migrator.Migrate()
	if err != nil {
		log.Fatalf("Can not migrate db schema: %v", err)
	}
}
