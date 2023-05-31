package router

import (
	"api_echo_modelo/src/controllers"

	"github.com/labstack/echo/v4"
)

func RotasUsuarios(e *echo.Echo) {
	e.POST("/usuarios", controllers.CriarUsuario)
	e.GET("/usuarios", controllers.BuscarUsuarios)
	e.GET("/usuarios/:usuarioId", controllers.BuscarUsuario)
	e.PUT("/usuarios/:usuarioId", controllers.AtualizarUsuario)
	e.DELETE("/usuarios/:usuarioId", controllers.DeletarUsuario)
}
