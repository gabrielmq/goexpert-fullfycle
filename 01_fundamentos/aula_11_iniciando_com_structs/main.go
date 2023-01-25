// Aula 11 - Iniciando com Structs
package main

import "fmt"

// exemplo de definção de uma struct
// struct representa uma estrutura de dados que representa um tipo de dado
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 28,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", gabriel.Nome, gabriel.Idade, gabriel.Ativo)

	// manipulando struct
	gabriel.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", gabriel.Nome, gabriel.Idade, gabriel.Ativo)
}
