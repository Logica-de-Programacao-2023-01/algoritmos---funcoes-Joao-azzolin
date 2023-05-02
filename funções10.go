// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com os valores ordenados de forma crescente. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
	"sort"
)

func crescente(slice []int) ([]int, error) {

	if len(slice) == 0 {

		return []int{0}, fmt.Errorf("o Slice está vazio")

	}

	sortedSlice := make([]int, len(slice))

	copy(sortedSlice, slice)

	sort.Ints(sortedSlice)

	return sortedSlice, nil

}

func main() {

	result, err := crescente([]int{1, 2, 7, 3, 4, 6, 5})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("Os valores de forma ordenada são: %d", result)

}
