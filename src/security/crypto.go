package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
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

func CriptografarTexto(textoClaro string, chave string) (string, error) {
	bloco, err := aes.NewCipher([]byte(chave))
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(bloco, iv)
	ciphertext := make([]byte, len(textoClaro))
	stream.XORKeyStream(ciphertext, []byte(textoClaro))

	ciphertext = append(iv, ciphertext...)
	return fmt.Sprintf("%x", ciphertext), nil
}

func DescriptografarTexto(textoCifrado string, chave string) (string, error) {
	bloco, err := aes.NewCipher([]byte(chave))
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(textoCifrado)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))

	stream := cipher.NewCFBDecrypter(bloco, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
