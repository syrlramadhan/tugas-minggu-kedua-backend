package repository

import (
	"context"
	"golang-database-user/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUserRepository_Success(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	UserRepository := NewUserRepositoryImpl(sql)

	users, err := UserRepository.ReadUsers(ctx)

	assert.NotNil(t, users)
	assert.Nil(t, err)
}

func TestReadUserRepository_Fail(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	UserRepository := NewUserRepositoryImpl(sql)

	_, err = UserRepository.ReadUsers(ctx)

	assert.NotNil(t, err)
}