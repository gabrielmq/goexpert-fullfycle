package main

import (
	"fmt"

	// necessario passar o nome completo do modulo + diretorio que quer importar as funcoes
	"github.com/gabrielmq/go-mod/math" // importando outro pacote do projeto
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
