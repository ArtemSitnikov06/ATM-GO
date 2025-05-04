package http

import (
	"ATMSystem/internal/BusinessLogic"
	"ATMSystem/internal/Presentation/http/Handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, atm BusinessLogic.AtmSystem) {
	h := Handlers.NewHandler(atm)

	e.POST("/api/users", h.CreateUser)
	e.GET("/api/users/:login", h.FindUserByLogin)
	e.DELETE("api/users/:login", h.DeleteUserByLogin)
	e.GET("/api/users", h.GetUsers)
	e.PUT("/api/users/login", h.ChangeLogin)
	e.POST("/api/bills", h.CreateBill)
	e.POST("/api/bills/deposit", h.Deposit)
	e.POST("/api/bills/withdraw", h.Withdraw)
}
