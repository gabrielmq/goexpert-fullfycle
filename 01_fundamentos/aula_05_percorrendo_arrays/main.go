// Aula 05 - Percorrendo Arrays
package main

import "fmt"

func main() {
	// array tem um tamanho fixo que pode ser acessado por indice
	// pesquisa por indice é mais rápida, inserção e remoção são mais demoradas
	// erros irao acontecer se tentar adicionar mais elementos que o array suporta
	// sempre começa pelo indice zero

	// definindo um array
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Printf("Elemento do array na posicao 2: %v\n", array[2])
	fmt.Printf("Elemento do array na ultima posicao: %v\n", array[len(array)-1])

	fmt.Println("imprimindo valores do array...")
	for indice, valor := range array {
		fmt.Println("valor", valor, "no indice", indice)
	}
	// for i := 0; i < len(array); i++ {
	// 	fmt.Println("imprimindo valor do array", array[i])
	// }
}
