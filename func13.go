package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	numero := 12345

	soma, err := somarDigitos(numero)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos dígitos:", soma)
	}
}

func somarDigitos(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Número negativo")
	}

	digitos := strconv.Itoa(num)
	soma := 0

	for _, d := range digitos {
		digito, _ := strconv.Atoi(string(d))
		soma += digito
	}

	return soma, nil
}
