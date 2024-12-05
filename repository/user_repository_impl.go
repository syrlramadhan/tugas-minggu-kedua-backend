package repository

import (
	"context"
	"database/sql"
	"fmt"

	// "errors"
	"golang-database-user/model"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repo *userRepositoryImpl) EmailExists(ctx context.Context, email string) (bool, error) {
	query := "SELECT COUNT(1) FROM mst_user WHERE email = $1 LIMIT 1"

	var count int
	err := repo.DB.QueryRowContext(ctx, query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// InsertUser : fungsi untuk melakukan query ke dalam database. ( contoh di bawah ini adalah fungsi untuk membuat data user )
func (repo *userRepositoryImpl) InsertUser(ctx context.Context, user model.MstUser) (model.MstUser, error) {

	query := "INSERT INTO mst_user(id_user, name, email, password, phone_number, role_id) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := repo.DB.ExecContext(ctx, query, user.IdUser, user.Name, user.Email, user.Password, user.PhoneNumber, user.Role.IdRole)
	if err != nil {
		return model.MstUser{}, err
	}

	return user, err
}

func (repo *userRepositoryImpl) UpdateUser(ctx context.Context, user model.MstUser, userId string) (model.MstUser, error) {
    var exists bool
    err := repo.DB.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM mst_user WHERE id_user = $1)", userId).Scan(&exists)
    if err != nil {
        return model.MstUser{}, fmt.Errorf("failed to check user existence: %w", err)
    }
    if !exists {
        return model.MstUser{}, fmt.Errorf("user with ID %s not found", userId)
    }

    query := "UPDATE mst_user SET name = $1, email = $2, password = $3, phone_number = $4 WHERE id_user = $5 RETURNING id_user, name, email, password, phone_number"

    row := repo.DB.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.PhoneNumber, userId)

    var updatedUser model.MstUser
    err = row.Scan(&updatedUser.IdUser, &updatedUser.Name, &updatedUser.Email, &updatedUser.Password, &updatedUser.PhoneNumber)
    if err != nil {
        return model.MstUser{}, fmt.Errorf("failed to scan updated user: %w", err)
    }

    return updatedUser, err
}


func (repo *userRepositoryImpl) ReadUsers(ctx context.Context) ([]model.MstUser, error) {
	query := "SELECT u.id_user, u.name, u.email, u.phone_number, r.role_name FROM mst_user u LEFT JOIN mst_role r ON u.role_id = r.id_role"

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed : %w", err)
	}
	defer rows.Close()

	var users []model.MstUser

	for rows.Next() {
		var user model.MstUser
		var role model.MstRole

		err := rows.Scan(&user.IdUser, &user.Name, &user.Email, &user.PhoneNumber, &role.RoleName)
		if err != nil {
			return nil, fmt.Errorf("failed : %w", err)
		}

		user.Role = role
		users = append(users, user)
	}

	return users, fmt.Errorf("failed : %w", err)
}


func (repo *userRepositoryImpl) DeleteUser(ctx context.Context, userId string) (model.MstUser, error) {
    var user model.MstUser

    querySelect := "SELECT id_user, name, email, phone_number FROM mst_user WHERE id_user = $1"
    err := repo.DB.QueryRowContext(ctx, querySelect, userId).Scan(&user.IdUser, &user.Name, &user.Email, &user.PhoneNumber)
    if err != nil {
        if err == sql.ErrNoRows {
            return model.MstUser{}, fmt.Errorf("user with ID %s not found", userId)
        }
        return model.MstUser{}, err
    }

    queryDelete := "DELETE FROM mst_user WHERE id_user = $1"
    _, err = repo.DB.ExecContext(ctx, queryDelete, userId)
    if err != nil {
        return model.MstUser{}, err
    }

    return user, nil
}
