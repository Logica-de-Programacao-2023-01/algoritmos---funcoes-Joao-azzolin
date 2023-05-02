
// Crie uma função que receba um número inteiro como parâmetro e retorne a soma de todos os seus dígitos. Caso o número seja negativo, retorne um erro.

package main

import (
	"fmt"
)

func somaDigitos(x int) (int, error) {

	if x < 0 {

		return 0, fmt.Errorf("o número é negativo")

	}

	soma := 0

	for x > 0 {

		soma += x % 10

		x /= 10

	}

	return soma, nil

}

func main() {

	result, err := somaDigitos(9876)

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("A soma dos digitos é: %d", result)

}
