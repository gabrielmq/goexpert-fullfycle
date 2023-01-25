// Aula 13 - Métodos em Structs
package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// exemplo de declarção de metodo; o () depois do func associa a função a struct Cliente
func (c Cliente) Desativar() {
	c.Ativo = false
}

func main() {
	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 28,
		Ativo: true,
	}

	fmt.Println("Desativando cliente", gabriel.Nome)

	gabriel.Desativar()
	fmt.Printf("Cliente %s foi desativado", gabriel.Nome)
}
