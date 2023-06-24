package main

import "fmt"

func main() {
	palavras := []string{"Olá", "mundo", "!", "Estou", "aqui"}

	resultado := concatenarStrings(palavras)

	fmt.Println("Concatenação:", resultado)
}

func concatenarStrings(strs []string) string {
	resultado := ""

	for _, str := range strs {
		resultado += str
	}

	return resultado
}
