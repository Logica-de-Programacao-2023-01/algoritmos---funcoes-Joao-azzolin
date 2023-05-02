// Crie uma função que receba uma string e retorne a quantidade de vogais presentes.

package main

import (
	"fmt"
	"strings"
)

func qntdvogais(s string) int {

	vogais := []string{"a", "e", "i", "o", "u"}

	soma := 0

	for i := 0; i < len(vogais); i++ {

		if strings.Contains(s, vogais[i]) {

			soma++

		}

	}

	return soma

}

func main() {

	result := qntdvogais("Hello, world!")

	fmt.Printf("A quantidade de vogais presente nessa String é: %d", result)

}
