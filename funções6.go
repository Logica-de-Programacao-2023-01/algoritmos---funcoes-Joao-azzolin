package main

import (
"fmt"
)

func concmaisvirgula(slice []string) (string, error) {

if len(slice) == 0 {

return "", fmt.Errorf("não tem nenhum texto para ser lido")

}

soma := slice[0]

for i := 1; i < len(slice); i++ {

soma += ", "

soma += slice[i]

}

return soma, nil

}

func main() {

result, err := concmaisvirgula([]string{"Hello", "world!"})

if err != nil {

fmt.Printf("Houve um erro ao executar o programa: %s.", err)

}

fmt.Printf("A concatenação desses elementos separados por vírgula é: %s", result)

}
