package service

import (
	"context"
	"github.com/google/uuid"
	"golang-database-user/model"
	"golang-database-user/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository, roleRepository repository.RoleRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		RoleRepository: roleRepository,
	}
}

// CreateUser : Fungsi untuk melakukan validasi dan logika pada program.
// contohnya jika anda di suruh untuk melakukan validasi untuk pengecekan nomor hp yang tidak boleh sama di dalam table mst_user
func (userService UserServiceImpl) CreateUser(ctx context.Context, userModel model.MstUser) model.MstUser {

	uuidUser := uuid.New().String()

	theRole, err := userService.RoleRepository.FindMstRole(ctx, "ROLE002")
	if err != nil {
		panic(err)
	}

	user := model.MstUser{
		IdUser:      uuidUser,
		Name:        userModel.Name,
		Email:       userModel.Email,
		Password:    userModel.Password,
		PhoneNumber: userModel.PhoneNumber,
		Role:        theRole,
	}

	insertUser, err := userService.UserRepository.InsertUser(ctx, user)
	if err != nil {
		panic(err)
	}

	return insertUser
}
