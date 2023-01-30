// Aula 05 - Busca CEP CLI
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// converte json para struct https://mholt.github.io/json-to-go/
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	filename := "cidades.txt"
	// pegando os argumentos da linha de comando com os.Args[1:]
	for _, cep := range os.Args[1:] {
		// realizando a requiscao
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			os.Exit(-1)
		}
		// fechando os recursos do body após tudo ser executado
		defer req.Body.Close()

		// lendo resposta do request
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			os.Exit(-1)
		}

		// convertendo o json da resposta para struct
		var data ViaCEP
		if err := json.Unmarshal(res, &data); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
			os.Exit(-1)
		}

		// criando um arquivo para armazenar os ceps buscados
		file, err := os.Create(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
			os.Exit(-1)
		}
		defer file.Close()

		// escrevendo no arquivo
		value := fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf)
		_, err = file.WriteString(value)
		fmt.Println("Arquivo criado com sucesso!")
	}
}
