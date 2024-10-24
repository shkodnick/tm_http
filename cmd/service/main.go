package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"app/internal/renderer"
	"app/internal/router"
)

func main() {
	e := echo.New()

	e.Renderer = renderer.NewRenderer("templates/*.html")

	router.RegisterRoutes(e)
	fmt.Println("localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}