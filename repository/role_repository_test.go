package repository

import (
	"context"
	"golang-database-user/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleRepository_Success(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	newRoleRepositoryImpl := NewRoleRepositoryImpl(sql)

	ctx := context.Background()

	role, err := newRoleRepositoryImpl.FindMstRole(ctx, "ROLE001")
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, role)
	assert.Nil(t, err)
	assert.Equal(t, "ROLE001", role.IdRole)
	assert.Equal(t, "Ketua Umum", role.RoleName)
}

func TestRoleRepository_Fail(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	newRoleRepositoryImpl := NewRoleRepositoryImpl(sql)

	ctx := context.Background()

	_, err = newRoleRepositoryImpl.FindMstRole(ctx, "ROLE009")

	assert.NotNil(t, err)
}