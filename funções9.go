// Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string. Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro.

package main

import (
	"fmt"
	"strings"
)

func palavras(s string) ([]string, error) {

	if len(s) == 0 {

		return []string{""}, fmt.Errorf("não tem nenhum texto para ser lido")

	}

	words := strings.Split(s, " ")

	return words, nil

}

func main() {

	result, err := palavras("Salve rapaziadinha da quebrada")

	if err != nil {

		fmt.Printf("Houve um erro ao executar o programa: %s.", err)

	}

	fmt.Printf("O resultado é: %s", result)

}
