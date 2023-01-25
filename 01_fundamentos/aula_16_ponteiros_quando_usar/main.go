// Aula 16 - Ponteiros, quando usar
package main

func sum(a, b int) int {
	return a + b
}

func sum2(a, b *int) int {
	*a = 50 // alterando na memoria os valores das variaveis
	*b = 50
	return *a + *b
}

// & -> referenciamento de memoria
// * -> desreferenciamento de memoria

// nao devemos usar ponteiros quando queremos apenas passar uma cópia dos dados
// devemos usar ponteiros quando queremos tornar os dados mutáveis em qualquer ponto do código, alterando valor direto a memoria

func main() {
	a := 10
	b := 20

	// vai ser passada uma cópia dos valores das variaveis por parametro na funçao sum
	println(sum(a, b))

	// vai ser passada os valores reais por referencia de endereço na memoria usando & para a funcao sum2
	println(sum2(&a, &b))
	println(a)
	println(b)
}
