package controllers

import (
	"api_echo_modelo/src/database"
	"api_echo_modelo/src/models"
	"api_echo_modelo/src/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CriarTeste cria um Teste no banco de dados.
func CriarTeste(c echo.Context) error {
	var test models.Teste
	err := c.Bind(&test)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = test.Validar()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.Conectar()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTeste(db)
	NovoTesteId, err := repositorio.CriarTeste(test)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, NovoTesteId)
}
