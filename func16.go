package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	texto := "Olá, mundo!"

	resultado, err := substituirVogaisPorAsterisco(texto)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func substituirVogaisPorAsterisco(texto string) (string, error) {
	if len(texto) == 0 {
		return "", errors.New("A string está vazia")
	}

	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	substituido := ""

	for _, char := range texto {
		caracter := string(char)

		if contains(vogais, caracter) {
			substituido += "*"
		} else {
			substituido += caracter
		}
	}

	return substituido, nil
}

func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
