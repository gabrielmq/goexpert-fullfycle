// 03 - Iterando channels com range
package main

import "fmt"

func main() {
	// criando um channel
	ch := make(chan int)

	// enviando dados para o channel em outra thread
	// sempre primeiro iniciar a thread que irá preencher o channel antes de comecar a ler
	go publish(ch)

	// lendo dados do channel na thread principal
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int) {
	// preenchendo o channel a cada iteração do for
	for i := 0; i < 10; i++ {
		// só vai ser enviada para o channel se ele estiver vazio
		// enquanto estiver cheio, o dado nao é enviado para o channel
		ch <- i
	}

	// indica que o channel nao vai receber mais nada nele
	// sempre fechar o channel quando ele não receber mais dados para evitar deadlock
	close(ch)
}
