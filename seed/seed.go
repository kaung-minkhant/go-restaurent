package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
	"github.com/kaung-minkhant/go-restaurent/seed/seedobjects"
	"github.com/kaung-minkhant/go-restaurent/utils"
)

func main() {
	// generateKey()
	// signature, err := rsa.SignPSS(rand.Reader)
	godotenv.Load("../.env")
	database.ConnectToDb()
	seed()
	fmt.Println("Seeding successful")
}

func sign() {
	msg := []byte("test message")

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)
	fmt.Println("message sum:", msgHashSum)
}

func generateKey() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// The public key is a part of the *rsa.PrivateKey struct
	publicKey := privateKey.PublicKey
	fmt.Printf("Public key: %v - Private key: %x\n", publicKey, privateKey.D.Bytes())
}

func seed() {
	// create role
	for _, role := range seedobjects.RolesToCreate {
		_, err := database.Db.CreateRole(context.Background(), role)
		if err != nil {
			fmt.Println("Seeding failed - creating roles", err)
			return
		}
	}

	// create permissions
	for _, permission := range seedobjects.PermissionsToCreate {
		_, err := database.Db.CreatePermission(context.Background(), models.CreatePermissionParams{
			Permission: permission.Permission,
			Method:     permission.Method,
			Route:      permission.Route,
		})
		if err != nil {
			fmt.Println("Seeding failed - creating permissions", err)
			return
		}
	}

	// crewate role permissions
	for _, rolePermission := range seedobjects.RolePermissionsToCreate {
		for _, permission := range rolePermission.Permissions {

			_, err := database.Db.CreateRolePermission(context.Background(), models.CreateRolePermissionParams{
				Role:       rolePermission.Role,
				Permission: permission,
			})
			if err != nil {
				fmt.Println("Seeding failed - creating role permissions", err)
				return
			}
		}
	}

	// create admin
	for _, user := range seedobjects.UsersToCreate {
		adminPass, err := utils.EncryptPassword(user.Password)
		if err != nil {
			fmt.Println("Seeding failed - creating hashed password", err)
			return
		}
		_, err = database.Db.CreateUser(context.Background(), models.CreateUserParams{
			ID:                uuid.New(),
			Name:              user.Name,
			EncryptedPassword: adminPass,
			Role:              user.Role,
			EmployeeID:        user.EmployeeID,
		})
		if err != nil {
			fmt.Println("Seeding failed - creating user", err)
			return
		}
	}
}
