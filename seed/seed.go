package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
	"github.com/kaung-minkhant/go-restaurent/seed/seedobjects"
	"github.com/kaung-minkhant/go-restaurent/utils"
)

func main() {

	database.ConnectToDb()
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
	fmt.Println("Seeding successful")
}
