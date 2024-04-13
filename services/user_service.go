package services

import (
	"crud-service/models"
	"crud-service/repositories"
	"errors"
)

type IUserService interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserById(userId string) ([]models.User, error)
	DeleteUserById(userId string) error
	updateUserById(userId string) (models.User, error)
}

type userService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) CreateUser(user models.User) (models.User, error) {

	if user.Username == " " || user.Password == "" {
		return models.User{}, errors.New("username and password required")
	}

	createUser, err := us.userRepository.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}

	return createUser, nil
}

func (us *userService) GetAllUser() ([]models.User, error) {

	getAllUser, err := us.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}

	return getAllUser, nil
}

func (us *userService) GetUserById(userId string) ([]models.User, error) {

	getUserById, err := us.userRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return getUserById, nil
}
func (us *userService) DeleteUserById(userId string) error {

	err := us.userRepository.DeleteUserById(userId)
	if err != nil {
		return err
	}

	return err
}
func (us *userService) UpdateUserById(userId string) (models.User, error) {

	updateUserById, err := us.userRepository.UpdateUserById(userId)
	if err != nil {
		return updateUserById, err
	}

	return updateUserById, err
}
