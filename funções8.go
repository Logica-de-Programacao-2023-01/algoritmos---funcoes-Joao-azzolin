
// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com apenas os números pares contidos no slice. Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func pares(slice []int) ([]int, error) {

	if len(slice) == 0 {

		return []int{0}, fmt.Errorf("a Slice está vazia")

	}

	var par []int

	for i := 0; i < len(slice); i++ {

		if slice[i]%2 == 0 {

			par = append(par, slice[i])

		}

	}

	return par, nil

}

func main() {

	result, err := pares([]int{1, 2, 3, 4, 5, 6, 8, 11})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("Os números pares dessa lista são: %d", result)

}
