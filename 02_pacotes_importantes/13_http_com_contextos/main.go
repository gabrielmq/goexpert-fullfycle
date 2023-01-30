// Aula 13 - HTTP com Contexts
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// criando um contexto vazio com Background()
	ctx := context.Background()

	// criando um novo contexto com 1s de timeout para ser cancelado
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	// funcao para cancelar o contexto após exceder o tempo de execucao ou o termino da execucao do contexto
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		"http://viacep.com.br/ws/02861030/json",
		nil,
	)
	if err != nil {
		panic(err)
	}

	// faz o Do com um client default
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
