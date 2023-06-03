package router

import (
	"api_echo_modelo/src/controllers"
	"api_echo_modelo/src/middlewares"

	"github.com/labstack/echo/v4"
)

func RotasUsuarios(e *echo.Echo) {
	// Rota sem middleware
	e.POST("/usuario/criar", controllers.CriarUsuario)

	// Grupo de rotas com middleware|
	grupoUsuario := e.Group("/usuarios")
	grupoUsuario.Use(middlewares.Autenticar)

	grupoUsuario.GET("/buscar", controllers.BuscarUsuarios)
	grupoUsuario.GET("/buscar/:usuarioId", controllers.BuscarUsuario)
	grupoUsuario.PUT("/atualizar/:usuarioId", controllers.AtualizarUsuario)
	grupoUsuario.DELETE("/deletar/:usuarioId", controllers.DeletarUsuario)
}
