package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kaung-minkhant/go-restaurent/database/models"
	_ "github.com/lib/pq"
)

var Db *models.Queries

func ConnectToDb() {
	connectionString := "postgres://root:root@127.0.0.1:5432/restaurant_management_system?sslmode=disable"
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
