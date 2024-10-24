package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"


	"app/internal/core/profile"
	"app/internal/handler/auth/jwt"
	"app/internal/models"
)

type RegisterHandler struct {
	ProfileSrv   *profile.ProfileService
}

func NewProfileHandler(ProfileSrv *profile.ProfileService) RegisterHandler {
	return RegisterHandler{
		ProfileSrv:   ProfileSrv,
	}
}

func (h *RegisterHandler) SignUpHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", nil)
}

func (h *RegisterHandler) SignUp(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	password2 := c.FormValue("password2")

	if password != password2 {
		return c.JSON(http.StatusInternalServerError, "Пароли не совподают")
	}

	exists, err := h.ProfileSrv.UserExists(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка проверки пользователя")
	}
	if exists {
		return c.JSON(http.StatusBadRequest, "Пользователь с таким email уже существует")
	}

	user := &models.SignUpParams{
		Email:    email,
		Password: password,
	}

	err = h.ProfileSrv.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка при создании пользователя")
	}

	token, err := jwt.GenerateToken(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка генерации токена")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Пользователь успешно зарегистрирован",
		"token":   token,
	})
}
