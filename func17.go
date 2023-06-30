package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func main() {
	slice := []string{"banana", "laranja", "abacaxi", "uva"}

	resultado, err := ordenarStrings(slice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	sort.Strings(slice)
	resultado := strings.Join(slice, ", ")

	return resultado, nil
}
