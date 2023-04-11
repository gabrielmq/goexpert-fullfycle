// 01 - Iniciando com Channels
package main

import "fmt"

// channels fazem a comunicação entre threads
func main() {
	// criando um channel vazio que vai receber strings
	ch := make(chan string)

	go func() {
		// adicionando uma string ao channel
		ch <- "Hello Channel!"
	}()

	// lendo e esvaziando o channel, imprimindo no console
	fmt.Println(<-ch)
}
