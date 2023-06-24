package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	numero := 17

	primo, err := verificarNumeroPrimo(numero)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("É primo:", primo)
	}
}

func verificarNumeroPrimo(num int) (bool, error) {
	if num < 0 {
		return false, errors.New("Número negativo")
	}

	if num < 2 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(num)))

	for i := 2; i <= limite; i++ {
		if num%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

