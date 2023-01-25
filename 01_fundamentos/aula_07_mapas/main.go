// Aula 07 - Mapas
package main

import "fmt"

func main() {
	// definição de um map já com valores-> map[tipo key]tipo valor
	// definicao de um map vazio -> x := map[string]int{} ou x := make(map[string]int)
	salarios := map[string]int{
		"Gabriel": 10,
	}

	// adiciona um novo elemento no map
	salarios["Teste"] = 1

	// imprime todo o map
	fmt.Println(salarios)

	// remove um elemento do map pela key
	delete(salarios, "Teste")

	// imprime todo o map
	fmt.Println(salarios)

	// imprime o valor de uma key do map
	fmt.Println(salarios["Gabriel"])

	// percorrendo o map
	for k, v := range salarios {
		fmt.Println("chave:", k, "valor:", v)
	}

	// percorrendo o map ignorando a key
	for _, v := range salarios {
		fmt.Println("valor:", v)
	}
}
