package main

import (
	"errors"
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	novosNumeros, err := filtrarNumerosPares(numeros)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Novos Números:", novosNumeros)
	}
}

func filtrarNumerosPares(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice vazio")
	}

	novosNumeros := make([]int, 0)

	for _, num := range nums {
		if num%2 == 0 {
			novosNumeros = append(novosNumeros, num)
		}
	}

	return novosNumeros, nil
}
