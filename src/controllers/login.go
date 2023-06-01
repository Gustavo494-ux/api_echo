package controllers

import (
	"api_echo_modelo/src/database"
	"api_echo_modelo/src/models"
	"api_echo_modelo/src/repository"
	"api_echo_modelo/src/security"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login
func Login(c echo.Context) error {
	var usuario models.Usuario
	erro := c.Bind(&usuario)
	if erro != nil {
		return c.String(http.StatusBadRequest, erro.Error())
	}

	db, erro := database.Conectar()
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuario(db)
	usuarioBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	if erro = security.CompararHash(usuarioBanco.Senha, usuario.Senha); erro != nil {
		return c.JSON(http.StatusInternalServerError, erro.Error())
	}

	return c.JSON(http.StatusOK, usuarioBanco)
}
