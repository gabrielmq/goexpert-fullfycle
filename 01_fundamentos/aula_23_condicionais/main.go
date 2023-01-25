// Aula 23 - Condicionais
package main

func main() {
	a := 1
	b := 2

	// exemplo de if
	if a > b {
		println(a)
	} else {
		println(b)
	}

	// exemplo de switch
	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default: // se nenhuma condicao atender os cases, o default é o padrao hehe
		println("infinito")
	}
}
