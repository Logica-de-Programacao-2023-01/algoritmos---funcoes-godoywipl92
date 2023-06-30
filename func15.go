package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 5
	slice := []int{1, 2, 3, 4, 5}

	presente, err := verificarPresencaNumero(numero, slice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Presente:", presente)
	}
}

func verificarPresencaNumero(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, valor := range slice {
		if valor == numero {
			return true, nil
		}
	}

	return false, nil
}
