// 08 - Channel com buffers
package main

func main() {
	// criando um channel bufferizado, que espera ser preenchido
	// com 5 informações mais de 1 vez
	// channels sem buffer, só é possivel colocar 1 msg por vez
	// channels bufferizados, é possivel colocar mais de 1 msg até o buffer ser enchido
	ch := make(chan string, 5)
	// conseguindo colocar mais msgs no channel
	ch <- "Hello"
	ch <- "World"

	// conseguindo ler mais msg do channel
	println(<-ch)
	println(<-ch)
}
