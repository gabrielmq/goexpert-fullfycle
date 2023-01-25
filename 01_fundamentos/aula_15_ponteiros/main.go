// Aula 15 - Ponteiros
package main

import "fmt"

func main() {
	// declarando uma variavel e criando um apontamento na memoria
	a := 10
	fmt.Println(a) // a variavel vai chamar o ponteiro que vai devolver o valor a ser printado

	// atribuindo a variavel ponteiro a referencia na memoria onde esta o valor de a, usando &
	// * e & representam o enderecamento da memoria
	var ponteiro *int = &a
	*ponteiro = 20         // atribuindo um novo valor no endereço da memoria
	fmt.Println(*ponteiro) // desreferenciando o ponteiro para obter o valor na memoria

	b := &a // b aponta para o endereco de memoria da variavel a
	*b = 30

	fmt.Println(*b)
	fmt.Println(a)
}
