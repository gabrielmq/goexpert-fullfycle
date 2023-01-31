// Aula 03 - Key/Value Context
package main

import (
	"context"
	"fmt"
)

func main() {
	// criando um contexto com metadados
	// context.WithValue() não é muito utilizado no dia a dia de desenv
	ctx := context.WithValue(context.Background(), "token", "pass")
	bookHotel(ctx)
}

// por convenção da linguagem o contexto sempre dever ser o primeiro parametro das funcoes
func bookHotel(ctx context.Context) {
	// obtendo dados dentro do contexto
	token := ctx.Value("token")
	fmt.Println(token)
}
