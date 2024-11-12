package controller

import (
	"context"
	"fmt"
	"golang-database-user/model"
	"golang-database-user/service"
)

func DefaultChoose() {
	fmt.Println("Incorrect Number")
}

func CreateUser(userService service.UserService) {

	ctx := context.Background()

	// buat inputan

	user := model.MstUser{
		Name:        "Supriadi Obo",
		Email:       "obo@gmail.com",
		Password:    "obotest",
		PhoneNumber: "731098419837",
	}

	mstUser := userService.CreateUser(ctx, user)

	fmt.Println(mstUser)
}
