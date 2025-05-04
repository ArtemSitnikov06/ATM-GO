package main

import (
	"ATMSystem/internal/BusinessLogic"
	"ATMSystem/internal/DataBase"
	"ATMSystem/internal/DataBase/Repositories"
	"ATMSystem/internal/Presentation/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := DataBase.DBinit()

	userRepo := Repositories.NewUserRepository(db)
	billRepo := Repositories.NewBillRepository(db)
	atm := BusinessLogic.NewAtm(userRepo, billRepo)

	e := echo.New()

	e.Use(middleware.CORS())

	http.RegisterRoutes(e, atm)

	e.Logger.Fatal(e.Start(":8080"))
}
