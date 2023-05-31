package main

import (
	"api_echo_modelo/src/configs"
	"api_echo_modelo/src/router"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.Carregar()
	e := router.Gerar()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from Echo & MySql"})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
