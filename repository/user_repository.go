package repository

import (
	"context"
	"golang-database-user/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user model.MstUser) (model.MstUser, error)
}
