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

	var nama, email, password, hp string
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan email : ")
	fmt.Scan(&email)
	fmt.Print("Masukkan password : ")
	fmt.Scan(&password)
	fmt.Print("Masukkan nomor hp : ")
	fmt.Scan(&hp)

	user := model.MstUser{
		Name:        nama,
		Email:       email,
		Password:    password,
		PhoneNumber: hp,
	}

	mstUser := userService.CreateUser(ctx, user)

	fmt.Println(mstUser)
}

func UpdateUser(userService service.UserService) {
	ctx := context.Background()

	var userId, nama, email, password, hp string
	fmt.Print("Masukkan id user yang ingin di update: ")
	fmt.Scanln(&userId)

	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan email : ")
	fmt.Scan(&email)
	fmt.Print("Masukkan password : ")
	fmt.Scan(&password)
	fmt.Print("Masukkan nomor hp : ")
	fmt.Scan(&hp)

	user := model.MstUser{
		Name:        nama,
		Email:       email,
		Password:    password,
		PhoneNumber: hp,
	}

	mstUser := userService.UpdateUser(ctx, user, userId)

	fmt.Println(mstUser)
}

func ReadUser(userService service.UserService) {
	ctx := context.Background()

	users, err := userService.ReadUsers(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nAll Users:")
	for _, mstUser := range users {
		fmt.Println("Id : ", mstUser.IdUser, "\nNama : ", mstUser.Name, "\nEmail : ", mstUser.Email, "\nNomor HP : ", mstUser.PhoneNumber)
		fmt.Println()
	}
}


func DeleteUser(userService service.UserService) {
	ctx := context.Background()

	var userId string
	fmt.Print("Masukkan id user yang ingin di hapus: ")
	fmt.Scanln(&userId)

	err := userService.DeleteUser(ctx, userId)
	if err != nil {
		fmt.Println("Gagal menghapus user:", err)
	} else {
		fmt.Println("User berhasil dihapus")
	}
}
