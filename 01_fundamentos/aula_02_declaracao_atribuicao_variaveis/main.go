// Aula 02 - Declaração e atribuição de variaveis
package main

// declara uma constante global
const constante = "constante"

// declara uma variavel global, por padrão se não foi atribuido valor o go vai inferir um valor default por de baixo dos panos
var a bool

// declara variaveis globais
var (
	b int
	c string = "teste"
	d float64
)

func main() {
	// declara variavel escopo local
	var e string

	// declara variavel escopo local e já atribui um valor
	f := "teste" // é feita a inferencia de tipos sem precisar explicitar o tipo e deve ser feita apenas na primeira atribuição
	var g string = "teste"

	println(constante) // constante global
	println(a)         // variavel global
	println(b)         // variavel global
	println(c)         // variavel global
	println(d)         // variavel global
	println(e)         // variavel local
	println(f)         // variavel local
	println(g)         // variavel local
}
