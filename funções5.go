// Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição do primeiro elemento igual ao valor no slice. Caso não encontre, retorne -1.

package main

import "fmt"

func posigualvalor(slice []int, x int) int {

	for i := 0; i < len(slice); i++ {

		if x == slice[i] {

			return x

		}

	}

	return -1

}

func main() {

	result := posigualvalor([]int{1, 2, 3, 4, 5}, 3)

	fmt.Printf("A posição do primeiro elemento igual ao valor escolhido é: %d", result)

}
