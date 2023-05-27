package routes

import "github.com/labstack/echo/v4"

func Gerar() *echo.Echo {
	e := echo.New()
	TesteRoute(e)
	return e
}
