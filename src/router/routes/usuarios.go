package router

import (
	"api_echo_modelo/src/controllers"
	"api_echo_modelo/src/middlewares"

	"github.com/labstack/echo/v4"
)

func RotasUsuarios(e *echo.Echo) {
	// Rota sem middleware
	e.POST("/usuario", controllers.CriarUsuario)

	// Grupo de rotas com middleware|
	grupoUsuario := e.Group("/usuarios")
	grupoUsuario.Use(middlewares.JWTAuthentication)

	grupoUsuario.GET("", controllers.BuscarUsuarios)
	grupoUsuario.GET("/:usuarioId", controllers.BuscarUsuario)
	grupoUsuario.PUT("/:usuarioId", controllers.AtualizarUsuario)
	grupoUsuario.DELETE("/:usuarioId", controllers.DeletarUsuario)
}
