// Escreva uma função que receba dois slices de inteiros como parâmetros e retorne um novo slice contendo apenas os números presentes nos dois slices. Caso um dos slices esteja vazio, retorne um erro.

package main

import (
	"fmt"
)

func intersect(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, fmt.Errorf("pelo menos um dos Slices está vazio")
	}

	result := []int{}
	m := make(map[int]bool)

	for _, num := range slice1 {
		m[num] = true
	}

	for _, num := range slice2 {
		if m[num] {
			result = append(result, num)
		}
	}

	return result, nil
}

func main() {

	resultado, err := intersect([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o prgrama: %s.", err)

	}

	fmt.Println(resultado)

}
