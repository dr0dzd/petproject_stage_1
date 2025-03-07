package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u User) (User, error)
	GetUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	UpdateUserByID(id uint, u User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(u User) (User, error) {
	if err := ur.db.Create(&u).Error; err != nil {
		return User{}, err
	}
	return u, nil
}

func (ur *userRepository) GetUsers() ([]User, error) {
	var AllUsers []User
	err := ur.db.Find(&AllUsers).Error
	return AllUsers, err
}

func (ur *userRepository) GetUserByID(id uint) (User, error) {
	var user User
	err := ur.db.First(&user, id).Error
	return user, err
}

func (ur *userRepository) UpdateUserByID(id uint, u User) (User, error) {
	var localUser User
	if finderr := ur.db.First(&localUser, id).Error; finderr != nil {
		return User{}, finderr
	}
	if uperr := ur.db.Model(&localUser).Updates(&u).Error; uperr != nil {
		return User{}, uperr
	}
	return localUser, nil
}

func (ur *userRepository) DeleteUserByID(id uint) error {
	err := ur.db.Unscoped().Delete(&User{}, id).Error
	return err
}
