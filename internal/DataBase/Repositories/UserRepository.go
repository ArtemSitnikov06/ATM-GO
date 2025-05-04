package Repositories

import (
	"ATMSystem/internal/DataBase/DTO"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *DTO.UserDTO) error

	FindUserByLogin(login string) (*DTO.UserDTO, error)

	DeleteUser(user *DTO.UserDTO) error
	UpdateUser(user *DTO.UserDTO) error

	GetUsers() (*[]DTO.UserDTO, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (ur *userRepo) Create(user *DTO.UserDTO) error {
	return ur.db.Create(user).Error
}

func (ur *userRepo) FindUserByLogin(login string) (*DTO.UserDTO, error) {
	user := DTO.UserDTO{}
	err := ur.db.Where("login = ?", login).First(&user).Error
	return &user, err
}

func (ur *userRepo) DeleteUser(user *DTO.UserDTO) error {

	err := ur.db.Delete(user).Error
	return err
}

func (ur *userRepo) UpdateUser(user *DTO.UserDTO) error {
	err := ur.db.Save(user).Error
	return err
}

func (ur *userRepo) GetUsers() (*[]DTO.UserDTO, error) {
	var users []DTO.UserDTO
	result := ur.db.Find(&users)
	return &users, result.Error
}
