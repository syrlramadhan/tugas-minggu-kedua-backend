package repository

import (
	"context"
	"golang-database-user/config"
	"golang-database-user/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUserRepository_Success(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}
	defer sql.Close()

	UserRepository := NewUserRepositoryImpl(sql)

	ctx := context.Background()

	existingUserID := "a038dd8f-2d29-41d9-9c4e-8e89f24a418a"

	user := model.MstUser{
		Name:        "Andi Arya Galang",
		Email:       "galang@coconut.or.id",
		Password:    "123",
		PhoneNumber: "0987654321",
	}

	updateUser, err := UserRepository.UpdateUser(ctx, user, existingUserID)

	assert.NotNil(t, updateUser, "Data pengguna yang diperbarui seharusnya tidak nil")
	assert.Nil(t, err, "Seharusnya tidak ada error saat update pengguna")
	assert.NotEmpty(t, updateUser.IdUser, "ID pengguna yang diperbarui harus ada (tidak kosong)")
}

func TestUpdateUserRepository_Fail(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}
	defer sql.Close()

	UserRepository := NewUserRepositoryImpl(sql)

	ctx := context.Background()

	existingUserID := ""

	user := model.MstUser{}

	updateUser, err := UserRepository.UpdateUser(ctx, user, existingUserID)

	assert.NotNil(t, err)
	assert.Empty(t, updateUser.IdUser)
}