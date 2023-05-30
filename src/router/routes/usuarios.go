package router

import (
	"api_echo_modelo/src/controllers"

	"github.com/labstack/echo/v4"
)

func RotasUsuarios(e *echo.Echo) {
	e.POST("/usuarios", controllers.CriarUsuario)
	e.GET("/usuarios/:usuarioId", controllers.BuscarUsuario)
	/*e.PUT("/teste/:testeId", controllers.AtualizarTeste)
	e.DELETE("/teste/:testeId", controllers.DeletarTeste)*/
}
