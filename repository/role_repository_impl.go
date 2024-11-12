package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database-user/model"
)

type roleRepositoryImpl struct {
	DB *sql.DB
}

func NewRoleRepositoryImpl(db *sql.DB) RoleRepository {
	return &roleRepositoryImpl{
		DB: db,
	}
}

func (roleRepository roleRepositoryImpl) FindMstRole(ctx context.Context, roleId string) (model.MstRole, error) {

	query := "SELECT id_role, role_name FROM mst_role WHERE id_role = $1 LIMIT 1"

	rows, err := roleRepository.DB.QueryContext(ctx, query, roleId)

	role := model.MstRole{}

	if err != nil {
		return role, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	if rows.Next() {
		err := rows.Scan(&role.IdRole, &role.RoleName)
		if err != nil {
			return model.MstRole{}, err
		}
		return role, nil
	} else {
		return model.MstRole{}, errors.New("role not found")
	}
}
