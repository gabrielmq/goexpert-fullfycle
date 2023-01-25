// Aula 09 - Funções variádicas
package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}

// exemplo de função variadicas, os ... indicam a passagem de parametros variaveis, de 0 a n
func sum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
