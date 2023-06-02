package router

import (
	"api_echo_modelo/src/controllers"
	"api_echo_modelo/src/middlewares"

	"github.com/labstack/echo/v4"
)

func RotasUsuarios(e *echo.Echo) {
	e.POST("/usuarios", controllers.CriarUsuario)
	grupoUsuario := e.Group("/usuarios", middlewares.Autenticar)

	grupoUsuario.GET("", controllers.BuscarUsuarios)
	grupoUsuario.GET("/:usuarioId", controllers.BuscarUsuario)
	grupoUsuario.PUT("/:usuarioId", controllers.AtualizarUsuario)
	grupoUsuario.DELETE("/:usuarioId", controllers.DeletarUsuario)
}
