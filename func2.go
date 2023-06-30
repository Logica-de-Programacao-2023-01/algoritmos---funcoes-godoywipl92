package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Ola, mundo!"

	qtdVogais := contarVogais(texto)

	fmt.Println("Quantidade de vogais:", qtdVogais)
}

func contarVogais(s string) int {
	vogais := []string{"a", "e", "i", "o", "u"}
	qtdVogais := 0

	lowercase := strings.ToLower(s)

	for _, char := range lowercase {
		if strings.ContainsAny(string(char), vogais) {
			qtdVogais++
		}
	}

	return qtdVogais
}
