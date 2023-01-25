// Aula 22 - For
package main

func main() {
	// exemplo do for tradicional
	for i := 0; i < 10; i++ {
		println(i)
	}

	// exemplo de for range, range retorna o indice e o valor
	n := []int{1, 2, 3}
	for k, v := range n {
		println(k, v)
	}

	// exemplo de for condicional, igual while/do while
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// exemplo de for infinito, sem condicao
	for {
		println("for infinito")
	}
}
