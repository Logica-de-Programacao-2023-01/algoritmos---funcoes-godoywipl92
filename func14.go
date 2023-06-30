package main

import (
	"errors"
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	resultado, err := intersecaoSlices(slice1, slice2)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func intersecaoSlices(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	intersecao := make([]int, 0)

	for _, num1 := range slice1 {
		for _, num2 := range slice2 {
			if num1 == num2 {
				intersecao = append(intersecao, num1)
				break
			}
		}
	}

	return intersecao, nil
}
