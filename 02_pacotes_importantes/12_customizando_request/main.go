// Aula 12 - Ccustomizando Request
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// configurando timeout para as requisições http
	client := http.Client{
		Timeout: time.Second,
	}

	// cria uma nova instancia para um request customizado
	req, err := http.NewRequest(
		"GET",
		"http://viacep.com.br/ws/02861030/json/",
		nil,
	)
	if err != nil {
		panic(err)
	}

	// adicionando um header no request
	req.Header.Set("Accept", "application/json")

	// realizando o request customizado com o Do
	res, err := client.Do(req)
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
