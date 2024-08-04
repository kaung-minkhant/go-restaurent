package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
	"github.com/kaung-minkhant/go-restaurent/utils"
)

func main() {
	database.ConnectToDb()
	_, err := database.Db.CreateRole(context.Background(), "admin")
	if err != nil {
		fmt.Println("Seeding Admin failed")
		return
	}

	adminPass, err := utils.EncryptPassword("admin")
	if err != nil {
		fmt.Println("Seeding Admin failed")
		return
	}
	_, err = database.Db.CreateUser(context.Background(), models.CreateUserParams{
		ID:                uuid.New(),
		Name:              "admin",
		EncryptedPassword: adminPass,
		Role:              "admin",
		EmployeeID:        "admin",
	})
	if err != nil {
		fmt.Println("Seeding Admin failed")
		return
	}
	fmt.Println("Seeding Admin successful")
}
