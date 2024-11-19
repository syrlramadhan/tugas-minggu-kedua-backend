package service

import (
	"context"
	"errors"
	"golang-database-user/model"
	"golang-database-user/repository"

	"github.com/google/uuid"
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
	emailExists, err := userService.UserRepository.EmailExists(ctx, userModel.Email)
	if err != nil {
		return model.MstUser{}
	}

	if emailExists {
		panic("Email sudah terdaftar")
	}

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

func (userService UserServiceImpl) UpdateUser(ctx context.Context, userModel model.MstUser, userId string) model.MstUser {
	user := model.MstUser{
		Name:        userModel.Name,
		Email:       userModel.Email,
		Password:    userModel.Password,
		PhoneNumber: userModel.PhoneNumber,
	}

	updateUser, err := userService.UserRepository.UpdateUser(ctx, user, userId)
	if err != nil {
		panic(err)
	}

	return updateUser
}

func (userService UserServiceImpl) ReadUsers(ctx context.Context) ([]model.MstUser, error) {
	users, err := userService.UserRepository.ReadUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService UserServiceImpl) DeleteUser(ctx context.Context, userId string) error {
	if userId == "" {
		return errors.New("ID user tidak boleh kosong")
	}

	err := userService.UserRepository.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}