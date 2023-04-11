// 02 - Forever
package main

import "fmt"

// conceito de forever é quando queremos criar um channel para segurar um processo, até que esse
// channel seja preenchido, e seu valor lido esvaziando o channel
// se o channel nunca for preenchido, isso pode causar um deadlock e finalizar a
// aplicação com erro
func main() {
	// criando um channel para segurar o processo
	forever := make(chan bool)

	// executando um processo em outra thread
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

		// evitando um deadlock preenchendo o channel forever
		forever <- true
	}()

	// esperando ficar cheio para esvaziar o channel, isso vai segurar o processo
	// se o channel não tiver valor, pode ocorrer deadlock
	<-forever
}
