package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/router"
)

func init() {
	database.ConnectToDb()
}

func main() {
	fmt.Println("Welcome")

	r := router.InitRouter()

	fmt.Println("Starting server at port :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Starting up Server failed with error", err)
	}
}
