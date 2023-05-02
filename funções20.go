
// Escreva uma função que receba um slice de strings como parâmetro e retorne um novo slice contendo apenas as strings que possuem mais de 5 caracteres. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"errors"
	"fmt"
)

func filterByLength(slice []string) ([]string, error) {

	if len(slice) == 0 {

		return nil, errors.New("Slice vazio")

	}

	filtered := make([]string, 0)

	for _, str := range slice {

		if len(str) > 5 {

			filtered = append(filtered, str)

		}

	}

	return filtered, nil

}

func main() {

	result, err := filterByLength([]string{"Hello", "world", "Olá", "mundo"})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Println(result)

}
