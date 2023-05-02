package main

import (
	"errors"
	"fmt"
	"strings"
)

func substituirVogaisPorAsterisco(s string) (string, error) {

	if len(s) == 0 {

		return "", errors.New("string vazia")

	}

	vogais := "aeiouAEIOU"

	var result strings.Builder

	for _, r := range s {

		if strings.ContainsRune(vogais, r) {

			result.WriteRune('*')

		} else {

			result.WriteRune(r)

		}

	}

	return result.String(), nil

}

func main() {

	result, err := substituirVogaisPorAsterisco("Rafael")

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Println(result)

}
