
// Crie uma função que receba um slice de strings como parâmetro e retorne uma nova string com as strings em ordem alfabética. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
	"sort"
	"strings"
)

func ordemAlfabetica(slice []string) (string, error) {

	if len(slice) == 0 {

		return "", fmt.Errorf("o slice está vazio")

	}

	sort.Strings(slice)

	return strings.Join(slice, " "), nil

}

func main() {

	slice := []string{"banana", "abacaxi", "laranja", "uva", "manga"}

	resultado, err := ordemAlfabetica(slice)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(resultado)

}
