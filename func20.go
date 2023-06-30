package main

import (
	"errors"
	"fmt"
)

func main() {
	strings := []string{"apple", "banana", "orange", "kiwi", "grapefruit"}

	resultado, err := filtrarPorTamanho(strings)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func filtrarPorTamanho(strings []string) ([]string, error) {
	if len(strings) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]string, 0)

	for _, str := range strings {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}

	return resultado, nil
}
