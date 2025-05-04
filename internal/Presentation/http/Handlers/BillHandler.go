package Handlers

import (
	"ATMSystem/internal/BusinessLogic/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) CreateBill(c echo.Context) error {
	var bill Models.Bill

	if err := c.Bind(&bill); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	user, err := h.atm.FindUserByLogin("")
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	if err := h.atm.CreateBill(&bill, user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, bill)
}

func (h *Handler) Deposit(c echo.Context) error {
	var operation Operation

	if err := c.Bind(&operation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	err := h.atm.Deposit(operation.BillID, operation.Amount)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "bill not found"})
	}
	return c.NoContent(http.StatusOK)

}

func (h *Handler) Withdraw(c echo.Context) error {
	var operation Operation

	if err := c.Bind(&operation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	err := h.atm.Withdraw(operation.BillID, operation.Amount)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "bill not found"})
	}
	return c.NoContent(http.StatusOK)

}

type Operation struct {
	BillID uint
	Amount float64
}
