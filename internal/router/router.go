package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"app/internal/handler/auth"
)



func RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/signup", auth.SignUp)
	e.GET("/signup", auth.SignupHandler)
	// e.GET("/login", loginHandler)
	// e.POST("/login", loginPostHandler)
	// e.POST("/login", auth.Login)
	// e.POST("/logout", auth.Logout)
	// e.POST("/change_pw", auth.ChangePassword)
}
