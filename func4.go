package main

import "fmt"

func main() {
	numeros := []int{5, 10, 2, 8, 15}

	segundoMaiorValor := segundoMaior(numeros)

	fmt.Println("Segundo Maior Valor:", segundoMaiorValor)
}

func segundoMaior(nums []int) int {
	maior := nums[0]
	segundoMaior := nums[0]

	for _, num := range nums {
		if num > maior {
			segundoMaior = maior
			maior = num
		} else if num > segundoMaior && num < maior {
			segundoMaior = num
		}
	}

	return segundoMaior
}
