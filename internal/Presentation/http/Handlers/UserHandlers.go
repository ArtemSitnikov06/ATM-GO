package Handlers

import (
	"ATMSystem/internal/BusinessLogic"
	"ATMSystem/internal/BusinessLogic/Models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	atm BusinessLogic.AtmSystem
}

func NewHandler(system BusinessLogic.AtmSystem) *Handler {
	return &Handler{atm: system}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var user Models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := h.atm.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) FindUserByLogin(c echo.Context) error {
	login := c.Param("login")

	user, err := h.atm.FindUserByLogin(login)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, user)

}

func (h *Handler) DeleteUserByLogin(c echo.Context) error {
	login := c.Param("login")

	user, err := h.atm.FindUserByLogin(login)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	err = h.atm.DeleteUser(user)
	if err == nil {
		return c.JSON(http.StatusOK, map[string]string{"succes": "user deleted"})
	}
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete user"})
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.atm.GetAllUsers()
	if err == nil {
		return c.JSON(http.StatusOK, users)
	}
	return c.JSON(http.StatusNotFound, users)
}

func (h *Handler) ChangeLogin(c echo.Context) error {
	var request ChangleLoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	user, err := h.atm.FindUserByLogin(request.OldLogin)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	err = h.atm.ChangeLogin(user, request.NewLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to change login"})
	}
	return c.JSON(http.StatusOK, map[string]string{"succes": "login changed"})
}

type ChangleLoginRequest struct {
	OldLogin string
	NewLogin string
}
