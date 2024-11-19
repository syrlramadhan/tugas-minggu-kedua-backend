package repository

import (
	"context"
	"database/sql"
	// "errors"
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

func (repo *userRepositoryImpl) UpdateUser(ctx context.Context, user model.MstUser, userId string) (model.MstUser, error) {
	query := "UPDATE mst_user SET name = $1, email = $2, password = $3, phone_number = $4 WHERE id_user = $5"

	_, err := repo.DB.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.PhoneNumber, userId)
	if err != nil {
		return model.MstUser{}, err
	}

	return user, nil
}

func (repo *userRepositoryImpl) ReadUsers(ctx context.Context) ([]model.MstUser, error) {
	query := "SELECT id_user, name, email, phone_number FROM mst_user"

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.MstUser

	for rows.Next() {
		var user model.MstUser
		var role model.MstRole

		err := rows.Scan(&user.IdUser, &user.Name, &user.Email, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}

		user.Role = role
		users = append(users, user)
	}

	return users, nil
}


func (repo *userRepositoryImpl) DeleteUser(ctx context.Context, userId string) error {
	query := "DELETE FROM mst_user WHERE id_user = $1"

	_, err := repo.DB.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}
	return nil
}