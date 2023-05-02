
// Escreva uma função que receba um slice de inteiros e uma função como parâmetros. A função deve aplicar a função recebida em cada elemento do slice e retornar a soma de todos os resultados. Caso o slice esteja vazio, retorne um erro.

package main

import (
"errors"
"fmt"
)

func applyAndSum(slice []int, f func(int) int) (int, error) {

if len(slice) == 0 {

return 0, errors.New("slice vazio")

}

sum := 0

for _, value := range slice {

sum += f(value)

}

return sum, nil

}

func main() {

slice := []int{1, 2, 3, 4, 5}

f := func(x int) int { return x * 2 }

result, err := applyAndSum(slice, f)

if err != nil {

fmt.Println(err)

} else {

fmt.Println(result)

}

}
