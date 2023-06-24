package main

import "fmt"

func main() {
	numeros := []int{5, 10, 15, 20, 25}

	media := mediaAritmetica(numeros)

	fmt.Println("Média:", media)
}

func mediaAritmetica(nums []int) float64 {
	total := 0

	for _, num := range nums {
		total += num
	}

	media := float64(total) / float64(len(nums))

	return media
}
