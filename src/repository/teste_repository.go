package repository

import (
	"api_echo_modelo/src/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Teste struct {
	db *sqlx.DB
}

// NovoRepositorioDeTeste cria um novo repositorio de teste
func NovoRepositorioDeTeste(db *sqlx.DB) *Teste {
	return &Teste{db}
}

// CriarTeste adiciona um novo teste no banco de dados.
func (repositorio Teste) CriarTeste(Teste models.Teste) (uint64, error) {
	statement, erro := repositorio.db.Exec(
		`INSERT INTO Tabela_Teste (nome) VALUES (?)`,
		Teste.Nome,
	)
	linhasAfetadas, err := statement.RowsAffected()
	if err != nil {
		return 0, err
	}
	if erro != nil {
		return 0, erro
	}

	if linhasAfetadas == 0 {
		return 0, errors.New("nenhuma linha foi afetada, verifique os dados passados")
	}

	usuarioID, err := statement.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(usuarioID), nil
}

// BuscarTestePorId
func (repositorio Teste) BuscarTestePorId(testeId uint64) (models.Teste, error) {
	teste := models.Teste{}
	erro := repositorio.db.Get(&teste, "SELECT id,nome FROM Tabela_Teste WHERE id = ?", testeId)
	if teste.Id == 0 {
		return models.Teste{}, errors.New("nenhum usuário foi encontrado, verifique os dados fornecidos")
	}
	if erro != nil {
		return models.Teste{}, erro
	}
	return teste, nil
}

// AtualizarTeste
func (repositorio Teste) AtualizarTeste(testeId uint64, teste models.Teste) error {
	statement, erro := repositorio.db.Exec(
		` UPDATE Tabela_Teste SET nome = ? 
			where id = ? `,
		teste.Nome,
		testeId,
	)

	linhasAfetadas, err := statement.RowsAffected()
	if err != nil {
		return err
	}
	if erro != nil {
		return erro
	}

	if linhasAfetadas == 0 {
		return errors.New("nenhuma linha foi afetada, verifique os dados passados")
	}

	return nil
}

// DeletarTeste
func (repositorio Teste) DeletarTeste(testeId uint64) error {
	statement, erro := repositorio.db.Exec(
		` DELETE FROM Tabela_Teste  
			where id = ? `,
		testeId,
	)

	linhasAfetadas, err := statement.RowsAffected()
	if err != nil {
		return err
	}
	if erro != nil {
		return erro
	}

	if linhasAfetadas == 0 {
		return errors.New("nenhuma linha foi afetada, verifique os dados passados")
	}

	return nil
}
