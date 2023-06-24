package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 20

	resultado, err := obterPrimos(numero)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func obterPrimos(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("O número deve ser positivo")
	}

	primos := make([]int, 0)

	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}
