package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	palavras := []string{"Maçã", "banana", "Abacaxi", "laranja"}

	resultado, err := filtrarStringsMaiusculas(palavras)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func filtrarStringsMaiusculas(palavras []string) (string, error) {
	if len(palavras) == 0 {
		return "", errors.New("Slice vazio")
	}

	var stringsMaiusculas []string

	for _, palavra := range palavras {
		if len(palavra) > 0 && unicode.IsUpper(rune(palavra[0])) {
			stringsMaiusculas = append(stringsMaiusculas, palavra)
		}
	}

	resultado := strings.Join(stringsMaiusculas, ", ")

	return resultado, nil
}
