package auth
// import (
// 	"github.com/labstack/echo/v4"
// )
// func Login(c echo.Context) error {
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")

// 	var user User
// 	err := db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return c.JSON(http.StatusUnauthorized, "Неверный email или пароль")
// 		}
// 		return c.JSON(http.StatusInternalServerError, "Ошибка при проверке пользователя")
// 	}

// 	// Сравнение паролей
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		return c.JSON(http.StatusUnauthorized, "Неверный email или пароль")
// 	}

// 	// Создание токена JWT
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
// 		Email: user.Email,
// 		ID:    user.ID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Токен действует 72 часа
// 		},
// 	})

// 	// Подписание токена
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Ошибка при создании токена")
// 	}

// 	// Возвращаем токен клиенту
// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": tokenString,
// 	})
// }