package middlewares

import (
	"api_echo_modelo/src/security"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Autenticar(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if erro := security.ValidarToken(c); erro != nil {
			return c.JSON(http.StatusUnauthorized, "o token informado é inválido")
		}
		err := next(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return nil
	}
}
