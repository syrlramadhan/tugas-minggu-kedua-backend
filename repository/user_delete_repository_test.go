package repository

import (
	"context"
	"golang-database-user/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUserRepository_Success(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	UserRepository := NewUserRepositoryImpl(sql)

	userId := "5e9da7eb-a098-433e-851a-94c82cf87fa3"

	if userId == "" {
		panic("Id user tidak boleh kosong")
	}

	deleteUser, err := UserRepository.DeleteUser(ctx, userId)
	
	assert.Nil(t, err)
	assert.NotNil(t, deleteUser)
}

func TestDeleteUserRepository_Fail(t *testing.T) {
    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }
    defer sql.Close()

    ctx := context.Background()

    UserRepository := NewUserRepositoryImpl(sql)

    userId := ""

    deleteUser, err := UserRepository.DeleteUser(ctx, userId)

    assert.NotNil(t, err)
    assert.Equal(t, "", deleteUser.IdUser)
}
