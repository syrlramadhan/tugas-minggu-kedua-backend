package main

import (
	"fmt"
	"golang-database-user/config"
	"golang-database-user/controller"
	"golang-database-user/repository"
	"golang-database-user/service"
)

func main() {
	fmt.Print("Golang Databases....")
	
	sql, err := config.OpenConnectionPostgresSQL()

	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepositoryImpl(sql)
	roleRepository := repository.NewRoleRepositoryImpl(sql)

	userService := service.NewUserServiceImpl(userRepository, roleRepository)

	for {
		fmt.Println()
		fmt.Println("1. Create User")
		fmt.Println("2. Update User")
		fmt.Println("3. Read User")
		fmt.Println("4. Delete User")
		fmt.Println("5. Exit")
		fmt.Println("------------------")
		fmt.Print("Select one of the numbers above : ")

		var scan int8
		fmt.Scanln(&scan)

		if scan == 5 {
			break
		}

		switch {
		case scan == 1:
			controller.CreateUser(userService)
		default:
			controller.DefaultChoose()
		}
	}
}
