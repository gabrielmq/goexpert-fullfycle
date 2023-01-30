// Aula 04 - Json
package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	// o nome nas tags são usados na conversao da struct para json
	Numero int `json:"numero"` // exemplo de tags em go
	Saldo  int `json:"saldo"`
	// Saldo  int `json:"-"` ignorando o valor de saldo no bind do json com a tag
}

func main() {
	conta := Conta{1, 100}

	// convertenda uma struct para json usando a funcao Marshal do pacote json
	// o retorno é em bytes, por isso é necessario fazer a conversao
	res, err := json.Marshal(conta) // Marshal é usado quando precisamos do retorno em uma variavel
	if err != nil {
		panic(err)
	}
	println(string(res))

	// converte para json já especificando o local de output do json
	// nesse formato nao temos o retorno do json em uma variavel
	if err := json.NewEncoder(os.Stdout).Encode(conta); err != nil {
		panic(err)
	}

	// convertendo json para struct usando json.Unmarshal() passando a referencia na memoria do dado a ser populado pelos valores do json
	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var contaX Conta
	if err := json.Unmarshal(jsonPuro, &contaX); err != nil {
		panic(err)
	}

	println(contaX.Saldo)
}
