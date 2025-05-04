package DTO

import "ATMSystem/internal/BusinessLogic/Models"

type UserDTO struct {
	ID        uint   `gorm:"primaryKey;unique"`
	Login     string `gorm:"unique"`
	Name      string
	Age       int
	HairColor Models.HairColor
	Gender    Models.Gender
}

func (UserDTO) TableName() string {
	return "users"
}
