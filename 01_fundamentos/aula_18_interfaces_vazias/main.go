// Aula 18 - Interfaces vazias
package main

import "fmt"

// exemplo de uma interface vazia que significa que essa interface implementa "todo mundo"
// a interface vazia aceita varios tipos
type Teste interface{}

func main() {
	var x Teste = 10
	var y Teste = "Hello World"
	showType(x)
	showType(y)
}

func showType(t Teste) {
	fmt.Printf("O tipo é do parametro é %T e o valor é %v\n", t, t)
}
