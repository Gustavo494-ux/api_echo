package security

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
)

func GerarHash(senhaDescriptografada string) (string, error) {
	h := sha512.New()
	_, err := h.Write([]byte(senhaDescriptografada))
	if err != nil {
		return "", err
	}
	senhaCriptografada := h.Sum(nil)
	hashHex := hex.EncodeToString(senhaCriptografada)
	return hashHex, nil
}

func CompararHash(senhaCriptografada, senhaDescriptografada string) error {
	hashDescriptografado, err := GerarHash(senhaDescriptografada)
	if err != nil {
		return err
	}
	senhaCriptografadaBytes := []byte(senhaCriptografada)

	if fmt.Sprintf("%x", hashDescriptografado) != fmt.Sprintf("%x", senhaCriptografadaBytes) {
		return errors.New("os dados de login estão são inválidos")
	}

	return nil
}
