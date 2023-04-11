// 01 - Iniciando com goroutines
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is runing\n", i, name)
		time.Sleep(time.Second)
	}
}

// main é a thread principal de execução de programas GO
func main() {
	// criando go rountines (threads) para executar a função task
	go task("A")
	go task("B")

	// executando uma função anonima em outra thread
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is runing\n", i, "anonima")
			time.Sleep(time.Second)
		}
	}()

	// sleep para segurar o processo da thread principal, pra ela nao morrer
	time.Sleep(30 * time.Second)
}
