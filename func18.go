package main

import (
	"errors"
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}

	resultado, err := aplicarFuncaoESomar(slice, quadrado)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func aplicarFuncaoESomar(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	soma := 0
	for _, valor := range slice {
		resultado := funcao(valor)
		soma += resultado
	}

	return soma, nil
}

func quadrado(numero int) int {
	return numero * numero
}
