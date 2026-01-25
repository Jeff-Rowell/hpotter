package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type FormData struct {
	Username string
	Password string
}

type LoginResponse struct {
	FormData FormData
	Error    string
}

func HandleLocalLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	formData := FormData{
		Username: username,
		Password: password,
	}

	response := &LoginResponse{
		FormData: formData,
		Error:    "invalid username or password",
	}
	return c.Render(http.StatusUnauthorized, "login-form", response)
}

func HandleLoginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "login-form", nil)
}
