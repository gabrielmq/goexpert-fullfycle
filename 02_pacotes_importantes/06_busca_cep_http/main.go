// Aula 06 - Busca CEP HTTP
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	// criando uma rota e configurando uma funcao para responder na rota
	http.HandleFunc("/", BuscaCepHandler)

	// iniciando um servidor HTTP
	http.ListenAndServe(":8080", nil)
}

// *http.Request representa as informacoes do request
// http.ResponseWriter enviar o response
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// adicionando no header do response o status code NOT_FOUND
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// pegando o query parameter do request
	param := r.URL.Query().Get("cep")
	if param == "" {
		// adicionando no header do response o status code BAD_REQUEST
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(param)
	if err != nil {
		// adicionando no header do response o status code INTERNAL_SERVER_ERROR
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// adicionando header Content-Type indicando que o retorno vai ser json
	w.Header().Set("Content-Type", "application/json")

	// adicionando no header do response o status code OK
	w.WriteHeader(http.StatusOK)

	// parseando a struct para json e retornando na api
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data ViaCEP
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
