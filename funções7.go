
package main

import (
	"fmt"
)

func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {

	if len(slice) == 0 {

		return []int{0}, fmt.Errorf("o Slice está vazio")

	}

	resultado := make([]int, len(slice))

	for i, valor := range slice {

		resultado[i] = funcao(valor)

	}

	return resultado, nil

}

func addum(valor int) int {

	return valor + 1

}

func main() {

	slice := []int{1, 2, 3, 4, 5}

	novoSlice, err := aplicarFuncao(slice, addum)

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("A resposta é: %d", novoSlice)

}
