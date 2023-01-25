// Aula 06 - Slices
package main

import "fmt"

func main() {
	// slice não tem tamanho fixo, ele pode crescer ou diminuir
	// por de baixo dos panos, o slice funciona com arrays
	// para lista de dados muito grande, ideal é criar um slice com uma capacidade proxima a quantidade de dados para nao prejudicar a performance do go por causa do redimencionamento do tamanho/capacidade do slice

	// estrutura de um slice
	// ponteiro -> faz o apontamento para o array interno do slice
	// tamanho -> diz até onde o slice deve cresce
	// capacidade -> diz a quantidade de dados que o slice consegue receber

	slice := []int{2, 4, 6, 8, 10, 12}

	// : -> é um ponto de corte no slice
	// slice[:indice] -> ignora tudo a direita a partir do indice especidicado
	// slice[indice:] -> ignora tudo a esquerda a partir do indice especidicado

	fmt.Printf("len=%d cap=%d | %v\n", len(slice), cap(slice), slice)

	// diminui o tamanho do slice removendo todos os elementos, matendo a mesma capacidade
	fmt.Printf("len=%d cap=%d | %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	// diminui o tamanho do slice ignorando todos os elementos a direita iniciando no indice 4, matendo a mesma capacidade
	fmt.Printf("len=%d cap=%d | %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	// diminui o tamanho e a capacidade do slice ignorando todos os elementos a esqueda iniciando no indice 2, matendo a mesma capacidade
	fmt.Printf("len=%d cap=%d | %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	// adicionando novos elementos a capacidade do slice aumenta, por causa do array interno
	slice = append(slice, 14)
	fmt.Printf("len=%d cap=%d | %v", len(slice[:2]), cap(slice[:2]), slice[:2])
}
