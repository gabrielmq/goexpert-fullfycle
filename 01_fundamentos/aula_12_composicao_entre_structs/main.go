// Aula 12 - Composição de Structs
package main

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // exemplo de composicao entre structs
	// Endereco Endereco usando a struct como um tipo na composicao
}

func main() {
	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 28,
		Ativo: true,
		Endereco: Endereco{
			Cidade: "Teste",
		},
	}

	// atribuicoes na struct endereco
	gabriel.Cidade = "São Paulo"
	gabriel.Endereco.Cidade = "São Paulo"
}
