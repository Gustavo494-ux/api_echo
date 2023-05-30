package repository

import (
	"api_echo_modelo/src/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Usuarios struct {
	db *sqlx.DB
}

// NovoRepositoDeUsuario cria um repositório de usuarios
func NovoRepositoDeUsuario(db *sqlx.DB) *Usuarios {
	return &Usuarios{db}
}

// CriarUsuario Adiciona um novo usuário
func (repositorio Usuarios) CriarUsuario(Usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Exec(
		` INSERT INTO Usuarios (nome, email, senha ) values (?,?,?) `,
		Usuario.Nome,
		Usuario.Email,
		Usuario.Senha,
	)
	linhasAfetadas, err := statement.RowsAffected()
	if linhasAfetadas == 0 {
		return 0, errors.New("nenhuma linha foi afetada, verifique os dados passados")
	}
	if err != nil {
		return 0, err
	}

	if erro != nil {
		return 0, erro
	}

	usuarioID, erro := statement.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(usuarioID), nil
}

// BuscarPorId busca um usuário pelo ID
func (repositorio Usuarios) BuscarPorId(usuarioId uint64) (models.Usuario, error) {
	usuarios := models.Usuario{}
	erro := repositorio.db.Get(&usuarios,
		` SELECT id,nome,email,criadoem FROM Usuarios WHERE id = ? `,
		usuarioId,
	)

	if usuarios.ID == 0 {
		return models.Usuario{}, errors.New("nenhum usuário foi encontrado, verifique os dados passados")
	}

	if erro != nil {
		return models.Usuario{}, erro
	}
	return usuarios, nil
}
