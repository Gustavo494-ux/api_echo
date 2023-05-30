package controllers

import (
	"api_echo_modelo/src/database"
	"api_echo_modelo/src/models"
	"api_echo_modelo/src/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(c echo.Context) error {
	var usuario models.Usuario
	erro := c.Bind(&usuario)
	if erro != nil {
		return c.String(http.StatusBadRequest, erro.Error())
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioID, erro := repositorio.CriarUsuario(usuario)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	usuario, erro = repositorio.BuscarPorId(usuarioID)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro)
	}

	return c.JSON(http.StatusCreated, usuario)
}
