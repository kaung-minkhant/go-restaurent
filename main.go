package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/router"
)

func init() {
	godotenv.Load(".env")
	sanityChecks()
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

func sanityChecks() {
	fmt.Println("Performing sanity checks")
	if dbUrl := os.Getenv("DB_URL"); dbUrl == "" {
		panic("DB_URL (database connection string) env required")
	}
	refreshLength := os.Getenv("REFRESH_TOKEN_LENGTH")
	if refreshLength == "" {
		panic("REFRESH_TOKEN_LENGTH env required")
	}
	if _, err := strconv.Atoi(refreshLength); err != nil {
		panic("REFRESH_TOKEN_LENGTH env invalid")
	}
	if jwt := os.Getenv("JWT_SECRET"); jwt == "" {
		panic("JWT_SECRET env invalid")
	}
	if tokenExpiry := os.Getenv("ACC_TOKEN_EXPIRY"); tokenExpiry == "" {
		panic("ACC_TOKEN_EXPIRY env required in minutes")
	}
	tokenExpiry := os.Getenv("ACC_TOKEN_EXPIRY")
	if _, err := strconv.Atoi(tokenExpiry); err != nil {
		panic("ACC_TOKEN_EXPIRY env invalid")
	}

	fmt.Println("Sanity checks done")
}
