package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	palavras := []string{"Olá", "mundo", "em", "Go"}

	resultado, err := concatenarComVirgulas(palavras)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Concatenação:", resultado)
	}
}

func concatenarComVirgulas(strs []string) (string, error) {
	if len(strs) == 0 {
		return "", errors.New("Slice vazio")
	}

	resultado := strings.Join(strs, ",")

	return resultado, nil
}
