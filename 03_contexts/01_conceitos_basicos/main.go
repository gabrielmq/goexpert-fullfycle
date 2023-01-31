// Aula 01 - Conceitos básicos
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// iniciando um contexto que roda em background na thread principal
	ctx := context.Background()

	// context.WithValue(...) cria um contexto passando valores de tempo
	// context.WithCancel(...) cancela o contexto a qualquer momento independente de um tempo
	// context.WithDeadline(...) cancela o contexto dado um determinado tempo
	// context.WithTimeout(...) cancela o contexto dado um tempo fazendo uma contagem regressiva

	// criando um contexto com timeout com base em um contexto já existente
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel() // sempre dar um cancel

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// select funciona como um switch, so que de forma assincrona aguardando resultado para tomar determinada acao
	select {
	// caso o contexto tenha sido finalizado por ter estourado o tempo de execucao
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout")
		return

	// executa caso esteja dentro do tempo de execucao do contexto
	case <-time.After(2 * time.Second):
		fmt.Println("Hotel booked.")
		return
	}
}
