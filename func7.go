package main

import (
	"errors"
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	novoSlice, err := aplicarFuncaoEmSlice(numeros, dobrar)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Novo Slice:", novoSlice)
	}
}

func aplicarFuncaoEmSlice(nums []int, fn func(int) int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice vazio")
	}

	novoSlice := make([]int, len(nums))

	for i, num := range nums {
		novoSlice[i] = fn(num)
	}

	return novoSlice, nil
}

func dobrar(num int) int {
	return num * 2
}
