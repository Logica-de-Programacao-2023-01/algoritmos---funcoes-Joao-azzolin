// Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele for um número primo e falso caso contrário. Caso o número seja negativo, retorne um erro.

package main

import (
	"fmt"
	"math"
)

func verprimo(x int) (bool, error) {

	if x < 0 {

		return false, fmt.Errorf("o número é negativo")

	}

	if x < 2 {

		return false, nil

	}

	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {

		if x%i == 0 {

			return false, nil

		}

	}

	return true, nil

}

func main() {

	result, err := verprimo(3)

	if err != nil {

		fmt.Printf("Houve um erro ao digitar um programa: %s.", err)

	}

	fmt.Println(result)

}
