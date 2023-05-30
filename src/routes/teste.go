package routes

import (
	"api_echo_modelo/src/controllers"

	"github.com/labstack/echo/v4"
)

func TesteRoute(e *echo.Echo) {
	e.POST("/teste", controllers.CriarTeste)
	e.GET("/teste/:testeId", controllers.BuscarTestePorId)
}
