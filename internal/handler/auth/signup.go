package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"app/internal/store"
	"app/internal/jwt"
	"app/internal/models"
)

func SignupHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", nil)
}

func SignUp(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	password2 := c.FormValue("password2")


	exists, err := store.UserExists(email)
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

	err = store.CreateUser(user)
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
