package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/kaung-minkhant/go-restaurent/database/models"
	_ "github.com/lib/pq"
)

var Db *models.Queries

func ConnectToDb() {

	connectionString := os.Getenv("DB_URL")
	fmt.Println("Opening db connection")
	dbInstance, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("cannot open db connection with error", err)
	}
	fmt.Println("Pinging db")
	if err := dbInstance.Ping(); err != nil {
		log.Fatal("Pinging the database failed with error", err)
	}
	fmt.Println("Pinging db successful")
	Db = models.New(dbInstance)
}
