package models

import (
	"errors"
	"strings"
)

type Teste struct {
	Id   uint64 `json:"id,omitempty" db:"id,omitempty"`
	Nome string `json:"nome" db:"nome,omitempty"`
}

func (t *Teste) Validar() error {
	switch {
	case strings.TrimSpace(t.Nome) == "":
		return errors.New("nome é um campo obrigatorio")
	}

	return nil
}
