package repositories

import (
	"crud-service/database"
	"crud-service/models"
)

type IUserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserById(userId string) ([]models.User, error)
	DeleteUserById(userId string) error
	UpdateUserById(userId string) (models.User, error)
}

type userRepository struct {
	db database.IDatabase
}

func NewUserRepository(db database.IDatabase) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user models.User) (models.User, error) {

	db := ur.db.GetDB()

	result := db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (ur *userRepository) GetAllUser() ([]models.User, error) {

	var user []models.User
	db := ur.db.GetDB()

	result := db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
func (ur *userRepository) GetUserById(userId string) ([]models.User, error) {

	var user []models.User
	db := ur.db.GetDB()

	result := db.First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *userRepository) DeleteUserById(userId string) error {

	var user models.User
	db := ur.db.GetDB()

	result := db.Delete(&user, userId)
	if result.Error != nil {
		return result.Error
	}

	return result.Error
}
func (ur *userRepository) UpdateUserById(userId string) (models.User, error) {

	var user models.User
	db := ur.db.GetDB()

	result := db.Delete(&user, userId)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return models.User{}, result.Error
}
