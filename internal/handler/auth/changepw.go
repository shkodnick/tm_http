package auth

// func ChangePassword(c echo.Context) error {
// 	email := c.FormValue("email")
// 	oldPassword := c.FormValue("old_password")
// 	newPassword := c.FormValue("new_password")

// 	var user User
// 	err := db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return c.JSON(http.StatusUnauthorized, "Неверный email или пароль")
// 		}
// 		return c.JSON(http.StatusInternalServerError, "Ошибка при проверке пользователя")
// 	} 

// 	// Проверка старого пароля
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
// 		return c.JSON(http.StatusUnauthorized, "Неверный старый пароль")
// 	}

// 	// Хеширование нового пароля
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Ошибка хеширования нового пароля")
// 	}

// 	// Обновление пароля в базе данных
// 	_, err = db.Exec("UPDATE users SET password=$1 WHERE email=$2", hashedPassword, email)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Ошибка при обновлении пароля")
// 	}

// 	return c.JSON(http.StatusOK, "Пароль успешно изменен")
// }