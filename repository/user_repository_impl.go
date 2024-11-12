package repository

import (
	"context"
	"database/sql"
	"golang-database-user/model"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

// InsertUser : fungsi untuk melakukan query ke dalam database. ( contoh di bawah ini adalah fungsi untuk membuat data user )
func (repo *userRepositoryImpl) InsertUser(ctx context.Context, user model.MstUser) (model.MstUser, error) {

	query := "INSERT INTO mst_user(id_user, name, email, password, phone_number, role_id) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := repo.DB.ExecContext(ctx, query, user.IdUser, user.Name, user.Email, user.Password, user.PhoneNumber, user.Role.IdRole)
	if err != nil {
		return model.MstUser{}, err
	}

	return user, nil
}
