package BusinessLogic

import (
	"ATMSystem/internal/BusinessLogic/Models"
	"ATMSystem/internal/DataBase/DTO"
)

func UserToDto(user *Models.User) DTO.UserDTO {
	return DTO.UserDTO{
		ID:        user.Id,
		Login:     user.Login,
		Name:      user.Name,
		Age:       user.Age,
		HairColor: user.HairColor,
		Gender:    user.Gender,
	}
}

func DtoToUser(user *DTO.UserDTO) Models.User {
	return Models.User{
		Id:        user.ID,
		Login:     user.Login,
		Name:      user.Name,
		Age:       user.Age,
		HairColor: user.HairColor,
		Gender:    user.Gender,
	}
}

func BillToDto(bill *Models.Bill, user *Models.User) DTO.BillDTO {
	return DTO.BillDTO{
		Balance: bill.Balance,
		Owner:   bill.Owner,
		User:    UserToDto(user),
	}
}

func DtoToBill(bill *DTO.BillDTO) Models.Bill {
	return Models.Bill{
		Id:      bill.ID,
		Balance: bill.Balance,
		Owner:   bill.Owner,
	}
}
