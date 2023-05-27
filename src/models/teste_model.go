package models

type Teste struct {
	Id   uint64 `json:"id,omitempty" db:"id,omitempty"`
	Nome string `json:"nome" db:"nome,omitempty"`
}
