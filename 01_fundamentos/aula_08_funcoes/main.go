// Aula 08 - Funções
package main

import (
	"errors"
	"fmt"
)

func main() {
	// atribuição de função com mais de um retorno
	v, err := sum(2, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}

// exemplo de função com mais de uma valor de retorno
func sum(a int, b int) (int, error) { // por convencao error deve ser o ultimo tipo do retorno
	sum := a + b
	if sum >= 50 {
		return 0, errors.New("Soma é maior que 50")
	}
	return sum, nil
}

// exemplo de declaração de função que recebe parametros e retorna apenas um valor
// func sum(a int, b int) int {
// 	return a + b
// }

// func sum(a, b int) int {
// 	return a + b
// }

// func sum(a, b int) { função void
// 	a + b
// }
