package Repositories

import (
	"ATMSystem/internal/DataBase/DTO"
	"gorm.io/gorm"
)

type billRepo struct {
	db *gorm.DB
}

type BillRepository interface {
	Create(bill *DTO.BillDTO) error
	Update(bill *DTO.BillDTO) error

	GetBillById(id uint) (*DTO.BillDTO, error)
}

func NewBillRepository(db *gorm.DB) BillRepository {

	return &billRepo{db: db}
}
func (br *billRepo) Create(bill *DTO.BillDTO) error {

	return br.db.Create(bill).Error
}

func (br *billRepo) Update(bill *DTO.BillDTO) error {
	return br.db.Save(bill).Error
}

func (br *billRepo) GetBillById(id uint) (*DTO.BillDTO, error) {
	var bill DTO.BillDTO
	result := br.db.First(&bill, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bill, nil
}
