package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	frase := "Esta é uma frase de exemplo"

	palavras, err := extrairPalavras(frase)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", palavras)
	}
}

func extrairPalavras(frase string) ([]string, error) {
	if frase == "" {
		return nil, errors.New("String vazia")
	}

	palavras := strings.Split(frase, " ")

	return palavras, nil
}
