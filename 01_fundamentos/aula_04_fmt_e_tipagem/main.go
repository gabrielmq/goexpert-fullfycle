// Aula 04 - Importando fmt e tipagem
package main

// importa o pacote fmt do GO, para printar no console
import "fmt"

type ID int

func main() {
	e := 1.2
	var f ID = 1
	// imprime uma msg com formatacao no console
	// %T (que representa o tipo) é um coringa que a função utiliza para substituir esse valor pelo parametro
	// %v representa o valor da variavel que sera usada
	fmt.Printf("tipo de E é %T", e)
	fmt.Printf("valor de F é %v", f)
}
