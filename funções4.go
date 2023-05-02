// Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.

package main

import "fmt"

func segmaior(slice []int) int {

	maior := slice[0]

	for i := 1; i < len(slice); i++ {

		if slice[i] > maior {

			maior = slice[i]

		}

	}

	segundo := slice[0]

	for j := 1; j < len(slice); j++ {

		if maior > slice[j] && slice[j] > segundo {

			segundo = slice[j]

		}

	}

	return segundo

}

func main() {

	result := segmaior([]int{6, 13, 5, 17, 9})

	fmt.Printf("O segundo maior número é: %d", result)

}
