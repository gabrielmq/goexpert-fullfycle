// Aula 19 - Type assertion
package main

import "fmt"

// esse padrao de type assertion é muito utilizado com interfaces vazias
// para sabermos os tipos dele
// o generics é um bom substituto para as interfaces vazias pois nao é
// necessário ficar fazer o type assertion
func main() {
	var a interface{} = "Gabriel"
	var b interface{} = 1
	println("convertendo (a.(string)) valor da variavel A para string:", a.(string)) // conversao direta para o tipo

	res, ok := b.(int) // convertendo com validacao e erro
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
