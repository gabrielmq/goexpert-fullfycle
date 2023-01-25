// Aula 10 - Closures (Funções anonimas)
package main

import "fmt"

func main() {
	// exemplo de declaração de função anonima/closure
	func() {
		fmt.Println("executando um função anonima sem retorno")
	}()

	func(s string) {
		fmt.Println("executando um função anonima", s)
	}("parametro funcao anonima")

	sum := func() int {
		return sum(1, 2, 3, 4, 5, 6) * 2
	}() // chamada da função imediatamente apos a definicao

	fmt.Println(sum)
}

func sum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
