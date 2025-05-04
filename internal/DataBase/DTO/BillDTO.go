package DTO

type BillDTO struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Balance float64
	Owner   uint
	User    UserDTO `gorm:"foreignKey:Owner"`
}

func (BillDTO) TableName() string {
	return "bills"
}
