package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	numeros := []int{5, 3, 1, 4, 2}

	novosNumeros, err := ordenarSlice(numeros)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Novos Números:", novosNumeros)
	}
}

func ordenarSlice(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice vazio")
	}

	novosNumeros := make([]int, len(nums))
	copy(novosNumeros, nums)

	sort.Ints(novosNumeros)

	return novosNumeros, nil
}
