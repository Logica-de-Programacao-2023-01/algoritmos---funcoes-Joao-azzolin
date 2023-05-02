// Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.

package main

import "fmt"

func conc(slings []string) string {

	soma := slings[0]

	for i := 0; i < len(slings)-1; i++ {

		soma = slings[i] + slings[i+1]

	}

	return soma

}

func main() {

	result := conc([]string{"Hello, ", "world!"})

	fmt.Printf("A concatenação do Slice de Strings é: %s", result)

}
