// Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.

package main

import "fmt"

func media(slice []int) int {

	soma := 0

	for _, x := range slice {

		soma += x

	}

	return soma / len(slice)

}

func main() {

	result := media([]int{1, 2, 3, 4, 5})

	fmt.Printf("A média é: %d.", result)

}
não consegui transferir do git bash
