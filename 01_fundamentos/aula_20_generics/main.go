// Aula 20 - Generics
package main

import "fmt"

// declarando uma constraint
// o ~ considera qualquer tipo que seja int ou float64
type Number interface {
	~int | ~float64
}

type MyNumber int

// exemplo de uma função generica
func Soma[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// func Soma[T int | float64](m map[string]T) T {
// 	var sum T
// 	for _, v := range m {
// 		sum += v
// 	}
// 	return sum
// }

// funcçoes genericas em Go só vao aceitar parametros do mesmo tipo, caso a funcao receba parametros genericos
// comparable é uma constraint para comparacao entre 2 valores
func Compara[T comparable](a T, b T) bool {
	return a == b
}

// func Compara[T any](a T, b T) bool {
// 	return a == b
// }

// func Compara[T Number](a T, b T) bool {
// 	return a == b
// }

func main() {
	m := map[string]int{
		"Gabriel": 1000,
		"Maria":   2000,
	}

	m2 := map[string]int{
		"Gabriel": 100.0,
		"Maria":   200.0,
	}

	m3 := map[string]MyNumber{
		"Gabriel": 100,
		"Maria":   200,
	}

	fmt.Println(Soma(m))
	fmt.Println(Soma(m2))
	fmt.Println(Soma(m3))

	fmt.Println(Compara(10, 10))
}
