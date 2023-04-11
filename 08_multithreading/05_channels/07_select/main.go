// 07 - Select, uma forma de escutar varios channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 4)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 4)
		ch2 <- 2
	}()

	// é uma forma de escutar mais de 1 channel
	// e o channel que retornar primeiro sera pego pelo select
	select {
	// se chegar msg primeiro aqui, sera printada a msg do channel ch1
	case msg1 := <-ch1:
		fmt.Println("received", msg1)

	// se chegar msg primeiro aqui, sera printada a msg do channel ch2
	case msg2 := <-ch2:
		fmt.Println("received", msg2)

	// define um timeout se não chegar msg em nenhum dos channels anteriores, para nao segurar o processo
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")

	// chega aqui quando nenhum dos channels ainda foi preenchido
	default:
		fmt.Println("not received")
	}
}
