package repository

import (
	"context"
	"golang-database-user/config"
	"golang-database-user/model"
	"testing"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertUserRepository_Success(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	UserRepository := NewUserRepositoryImpl(sql)
	RoleRepository := NewRoleRepositoryImpl(sql)

	uuidUser := uuid.New().String()

	theRole, err := RoleRepository.FindMstRole(ctx, "ROLE001")
	if err != nil {
		panic(err)
	}

	user := model.MstUser{
		IdUser:      uuidUser,
		Name:        "Rama",
		Email:       "rama@coconut.or.id",
		Password:    "123",
		PhoneNumber: "0987654321",
		Role:        theRole,
	}

	insertUser, err := UserRepository.InsertUser(ctx, user)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, insertUser)
	assert.Nil(t, err)
}

func TestInsertUserRepository_EmailExists(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	UserRepository := NewUserRepositoryImpl(sql)

	emailExists, _ := UserRepository.EmailExists(context.Background(), "syahrul@coconut.or.id")

    assert.True(t, emailExists, "Email Sudah Terdaftar")
}

func TestInsertUserRepository_Fail(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	UserRepository := NewUserRepositoryImpl(sql)

	user := model.MstUser{}

	_, err = UserRepository.InsertUser(ctx, user)

	assert.NotNil(t, err)
}