package service

import (
	"context"
	"golang-database-user/model"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.MstUser) model.MstUser
}
