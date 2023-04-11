// 04 - Range e Wait Groups
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}

	// indica para o wait group que serão executadas 10 tarefas em backgroud
	wg.Add(10)

	go publish(ch)

	go func() {
		for x := range ch {
			fmt.Printf("Received %d\n", x)
			// marca a tarefa como finalizada
			wg.Done()
		}
	}()

	// espera até que todas as tarefas do wait group sejam concluídas
	wg.Wait()
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
