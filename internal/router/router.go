package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"app/internal/bootstrap"
	"app/internal/handler/auth"
)

func RegisterRoutes(e *echo.Echo) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	profileSrv, err := bootstrap.CreateUserService()
	if err != nil {
		return err
	}

	profileHandler := auth.NewProfileHandler(profileSrv)

	e.POST("/signup", profileHandler.SignUp)
	e.GET("/signup", profileHandler.SignUpHandler)
	// e.GET("/login", loginHandler)
	// e.POST("/login", loginPostHandler)
	// e.POST("/login", auth.Login)
	// e.POST("/logout", auth.Logout)
	// e.POST("/change_pw", auth.ChangePassword)

	return nil
}
