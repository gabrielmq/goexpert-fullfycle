// Aula 11 - HttpClient POST
package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// configurando timeout para as requisições http
	client := http.Client{
		Timeout: time.Second,
	}

	// o body do POST precisa ser um slice de bytes bufferizado
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Teste"}`))

	// exemplo de requisição POST
	res, err := client.Post(
		"http://viacep.com.br/ws/02861030/json/",
		"application/json",
		jsonVar,
	)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// copia os dados do body e joga no STDOUT do sistema
	io.CopyBuffer(os.Stdout, res.Body, nil)
}
