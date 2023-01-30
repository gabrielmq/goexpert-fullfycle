// Aula 10 - HttpClient com Timeout
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

	res, err := client.Get("http://viacep.com.br/ws/02861030/json/")
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
