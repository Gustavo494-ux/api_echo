package router

import (
	router "api_echo_modelo/src/router/routes"

	"github.com/labstack/echo/v4"
)

func Gerar() *echo.Echo {
	e := echo.New()
	router.TesteRoute(e)
	router.RotasUsuarios(e)
	router.RotasLogin(e)
	return e
}
