package controllers

import (
	"api_echo_modelo/src/database"
	"api_echo_modelo/src/models"
	"api_echo_modelo/src/repository"
	"errors"
	"net/http"
	"strconv"

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

	testeBanco, err := repositorio.BuscarTestePorId(NovoTesteId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, testeBanco)
}

// BuscarTestePorId retorna os dados de um teste utilizando seu id para realizar a busca
func BuscarTestePorId(c echo.Context) error {
	testeId, erro := strconv.ParseUint(c.Param("testeId"), 10, 64)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTeste(db)
	testeBanco, err := repositorio.BuscarTestePorId(testeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, testeBanco)
}

// AtualizarTeste atualiza os dados do teste
func AtualizarTeste(c echo.Context) error {
	testeId, erro := strconv.ParseUint(c.Param("testeId"), 10, 64)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	var test models.Teste
	err := c.Bind(&test)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	test.Validar()

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTeste(db)
	testeBanco, err := repositorio.BuscarTestePorId(testeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if testeBanco.Id == 0 {
		return c.JSON(http.StatusNotFound, errors.New("teste não encontrado"))
	}

	if err = repositorio.AtualizarTeste(testeId, test); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	testeBanco, err = repositorio.BuscarTestePorId(testeId)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, testeBanco)
}

// DeletarTeste excluir os dados de um teste
func DeletarTeste(c echo.Context) error {
	testeId, erro := strconv.ParseUint(c.Param("testeId"), 10, 64)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, erro)
	}

	db, err := database.Conectar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTeste(db)
	testeBanco, err := repositorio.BuscarTestePorId(testeId)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if testeBanco.Id == 0 {
		return c.JSON(http.StatusNotFound, errors.New("teste não encontrado"))
	}

	if err = repositorio.DeletarTeste(testeId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, "registro excluido")
}
