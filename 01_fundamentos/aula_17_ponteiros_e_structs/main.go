// Aula 17 - Ponteiros e Structs
package main

import "fmt"

type Conta struct {
	saldo int
}

// Criando uma Conta retornando o apontamento na memoria para refletir globalmente onde a Conta for passada
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

// metodo sem ponteiro
// func (c Conta) simular(valor int) int {
// 	// propriedade saldo foi atualizada apenas no escopo do metodo simular, pois o valor na memoria nao foi de fato atualizado
// 	c.saldo += valor
// 	return c.saldo
// }

// metodo com referencia de memoria onde a Conta esta armazenada
func (c *Conta) simular(valor int) int {
	// propriedade saldo foi atualizada num escopo global pois seu valor foi atualizado direto no local da memoria onde ela estava armazenada
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := Conta{100}
	conta.simular(200)

	fmt.Printf("O valor do saldo %v\n", conta.saldo)
	fmt.Println(NewConta())
}
