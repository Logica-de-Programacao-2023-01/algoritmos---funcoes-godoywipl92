package main

import "fmt"

func main() {
	numeros := []int{5, 10, 2, 8, 15}
	valor := 8

	posicao := encontrarPosicao(numeros, valor)

	if posicao != -1 {
		fmt.Println("Posição:", posicao)
	} else {
		fmt.Println("Elemento não encontrado.")
	}
}

func encontrarPosicao(nums []int, valor int) int {
	for i, num := range nums {
		if num == valor {
			return i
		}
	}

	return -1
}
