
// Crie uma função que receba um número inteiro como parâmetro e retorne um novo slice contendo todos os números primos menores ou iguais a ele. Caso o número seja negativo, retorne um erro.

package main

import (
	"fmt"
	"math"
)

func main() {

	primes, err := getPrimes(20)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(primes)

}

func getPrimes(n int) ([]int, error) {

	if n < 0 {

		return nil, fmt.Errorf("número inválido: %d", n)

	}

	var primes []int

	for i := 2; i <= n; i++ {

		if isPrime(i) {

			primes = append(primes, i)

		}

	}

	return primes, nil

}

func isPrime(n int) bool {

	if n == 2 || n == 3 {

		return true

	}

	if n == 1 || n%2 == 0 {

		return false

	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {

		if n%i == 0 {

			return false

		}

	}

	return true

}
