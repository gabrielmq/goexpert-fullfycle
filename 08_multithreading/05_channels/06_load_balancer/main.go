// 06 - Criando um Load Balancer
// Load Balancer é uma forma de balancearmos uma carga de trabalho
// entre workers, para conseguirmos ter um processamento de forma concorrente e paralela
package main

import (
	"fmt"
	"time"
)

// 1 - iniciar a função que vai processar os dados do channel em novas threads
// 2 - iniciar o preenchimento do channel
func main() {
	ch := make(chan int)

	// inicializa os workers para ficarem processando o channel
	workers := 10
	for i := 1; i <= workers; i++ {
		go worker(i, ch)
	}

	// faz o envio de dados para o channel que a função worker le
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

// função responsável por realizar o processamento do channel dentro do conceito de LB
func worker(workerId int, ch <-chan int) {
	for data := range ch {
		fmt.Printf("Worker %d received %d\n", workerId, data)
		time.Sleep(time.Second)
	}
}
