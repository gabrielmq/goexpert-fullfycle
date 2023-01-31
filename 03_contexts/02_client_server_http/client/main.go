// Aula 02_2 - Contexts com client HTTP
package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// criando um contexto com timeout de 10s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// cancela o contexto se exceder os 10s
	defer cancel()

	// criando um request customizado com o context de 10s
	// se a  req passar de 10s o contexto é cancelado e a operacao é finalizada
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/", nil)
	if err != nil {
		panic(err)
	}

	// realizando a request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// copiando o body para ser impresso no terminal
	io.Copy(os.Stdout, res.Body)
}
