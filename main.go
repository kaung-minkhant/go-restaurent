package main

import (
	"fmt"

	"github.com/kaung-minkhant/go-restaurent/database"
)

func init() {
	database.ConnectToDb()
}

func main() {
	fmt.Println("Welcome")
}
