package BusinessLogic

import (
	"ATMSystem/internal/BusinessLogic/Models"
	"ATMSystem/internal/DataBase/Repositories"
	"errors"
)

type AtmSystem interface {
	CreateUser(user *Models.User) error

	FindUserByLogin(login string) (*Models.User, error)

	DeleteUser(user *Models.User) error

	ChangeLogin(user *Models.User, login string) error

	GetAllUsers() (*[]Models.User, error)

	CreateBill(bill *Models.Bill, user *Models.User) error

	Deposit(id uint, amount float64) error

	Withdraw(id uint, amount float64) error
}

type Atm struct {
	userRepo Repositories.UserRepository
	billRepo Repositories.BillRepository
}

func NewAtm(userRepo Repositories.UserRepository, billRepo Repositories.BillRepository) *Atm {
	return &Atm{
		userRepo: userRepo,
		billRepo: billRepo,
	}
}

func (atm *Atm) CreateUser(user *Models.User) error {
	userDTO := UserToDto(user)
	err := atm.userRepo.Create(&userDTO)
	if err == nil {
		user.Id = userDTO.ID
	}
	return err
}

func (atm *Atm) FindUserByLogin(login string) (*Models.User, error) {

	userDTO, err := atm.userRepo.FindUserByLogin(login)
	if err != nil {
		return nil, err
	}

	user := DtoToUser(userDTO)

	return &user, nil

}

func (atm *Atm) DeleteUser(user *Models.User) error {
	userDTO := UserToDto(user)
	err := atm.userRepo.DeleteUser(&userDTO)
	return err
}

func (atm *Atm) ChangeLogin(user *Models.User, login string) error {
	userDTO := UserToDto(user)
	userDTO.Login = login
	err := atm.userRepo.UpdateUser(&userDTO)
	return err
}

func (atm *Atm) GetAllUsers() (*[]Models.User, error) {
	users, err := atm.userRepo.GetUsers()
	var result []Models.User
	if err != nil {
		return &result, err
	}
	for _, u := range *users {
		user := DtoToUser(&u)
		result = append(result, user)
	}
	return &result, err

}

func (atm *Atm) CreateBill(bill *Models.Bill, user *Models.User) error {
	billDTO := BillToDto(bill, user)
	err := atm.billRepo.Create(&billDTO)
	if err == nil {
		bill.Id = billDTO.ID
	}
	return err
}

func (atm *Atm) Deposit(id uint, amount float64) error {
	billDTO, err := atm.billRepo.GetBillById(id)
	if billDTO != nil {
		billDTO.Balance += amount
		err1 := atm.billRepo.Update(billDTO)
		if err1 != nil {
			return err1
		}
	}
	return err
}

func (atm *Atm) Withdraw(id uint, amount float64) error {
	billDTO, err := atm.billRepo.GetBillById(id)
	if billDTO != nil {
		if billDTO.Balance >= amount {
			billDTO.Balance -= amount
			err1 := atm.billRepo.Update(billDTO)
			if err1 != nil {
				return err1
			}
		}

		return errors.New("not enough money in the account")
	}
	return err
}
