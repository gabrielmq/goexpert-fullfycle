// Aula 02 - Chamadas HTTP
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// exemplo de chamada http simples GET usando a funcao Get e o pacote http
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	// le o retorno da requisicao com a funcao ReadAll do pacote io
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// imprimindo o response convertendo de bytes para string
	fmt.Println(string(body))

	// fechando o stream de dados do response para nao ter leak de recursos
	req.Body.Close()
}
