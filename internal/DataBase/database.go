package DataBase

import (
	"ATMSystem/internal/DataBase/DTO"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBinit() *gorm.DB {
	var db *gorm.DB
	dsn := "host=localhost user=postgres password=postgres dbname=AtmDB port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&DTO.UserDTO{}, &DTO.BillDTO{})

	return db

}
