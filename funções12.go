
// Escreva uma função que receba um slice de strings como parâmetro e retorne uma string contendo apenas as strings que começam com uma letra maiúscula. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func apenasmaiusculas(slice []string) (string, error) {

	if len(slice) == 0 {

		return "", fmt.Errorf("o Slice está vazio")

	}

	var filtered []string

	for _, str := range slice {

		if len(str) > 0 && unicode.IsUpper([]rune(str)[0]) {

			filtered = append(filtered, str)

		}

	}

	if len(filtered) == 0 {

		return "", fmt.Errorf("nenhuma string com letra maiúscula encontrada")

	}

	return strings.Join(filtered, " "), nil

}

func main() {

	result, err := apenasmaiusculas([]string{"Hello", "world"})

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("As palavras que começam apenas com letra maiúscula são: %s", result)

}
