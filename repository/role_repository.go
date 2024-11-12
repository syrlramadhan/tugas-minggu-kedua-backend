package repository

import (
	"context"
	"golang-database-user/model"
)

type RoleRepository interface {
	FindMstRole(ctx context.Context, roleId string) (model.MstRole, error)
}
