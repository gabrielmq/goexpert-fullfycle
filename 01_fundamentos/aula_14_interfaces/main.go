// Aula 14 - Interfaces
package main

import "fmt"

// exemplo de declaração de interface
// Só podemos ter assinaturas de metodos nas interfaces do GO
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// automagicamente o GO entende por esse metodo que a struct Cliente esta implementando a interface Pessoa
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

	Desativacao(gabriel)
	fmt.Printf("Cliente %s foi desativado", gabriel.Nome)
}

// passando interface como um tipo de parametro para garantir polimorfismo
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}
