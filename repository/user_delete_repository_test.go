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
	
	assert.Nil(t, err, "Seharusnya tudakada error pada saat menghapus pengguna")
	assert.NotNil(t, deleteUser, "Data pengguna yang dihapus seharusnya tidak nil")
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

    assert.NotNil(t, err, "Seharusnya terjadi error pada saat menghapus pengguna")
    assert.Equal(t, "", deleteUser.IdUser, "Id pengguna yang ingin di hapus seharusnya nil")
}
