// Aula 03 - Defer
package main

func main() {
	println("primeira linha")
	// defer vai fazer essa linha ser executada por ultimo, depois que todas as instrucoes tiverem sido executada dentro da funcao
	defer println("segunda linha")
	println("terceira linha")
}
