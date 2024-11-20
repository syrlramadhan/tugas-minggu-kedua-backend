package controller

import (
	"bufio"
	"context"
	"fmt"
	"golang-database-user/model"
	"golang-database-user/service"
	"os"
	"strings"
)

func DefaultChoose() {
	fmt.Println("Incorrect Number")
}

func CreateUser(userService service.UserService) {
	reader := bufio.NewReader(os.Stdin)

	ctx := context.Background()

	fmt.Print("Masukkan nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Masukkan password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Masukkan nomor HP: ")
	hp, _ := reader.ReadString('\n')
	hp = strings.TrimSpace(hp)

	fmt.Println("\n"+`ROLE: (Keterangan : IdRole = Role Name)
ROLE001 = Ketua Umum
ROLE002 = Sekretaris Umum
ROLE003 = Bendahara Umum
ROLE004 = Departement Keorganisasian
ROLE005 = Departement Pembelajaran
ROLE006 = Departement Humas`)

	fmt.Print("\nMasukkan id role: ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	user := model.MstUser{
		Name:        nama,
		Email:       email,
		Password:    password,
		PhoneNumber: hp,
	}

	mstUser := userService.CreateUser(ctx, user, role)

	fmt.Println(mstUser)
}

func UpdateUser(userService service.UserService) {
	reader := bufio.NewReader(os.Stdin)

	ctx := context.Background()

	var userId string
	fmt.Print("Masukkan id user yang ingin di update: ")
	fmt.Scanln(&userId)

	fmt.Print("Masukkan nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Masukkan password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Masukkan nomor HP: ")
	hp, _ := reader.ReadString('\n')
	hp = strings.TrimSpace(hp)

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
		fmt.Println("Id : ", mstUser.IdUser, "\nNama : ", mstUser.Name, "\nEmail : ", mstUser.Email, "\nNomor HP : ", mstUser.PhoneNumber, "\nJabatan: ", mstUser.Role.RoleName)
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
