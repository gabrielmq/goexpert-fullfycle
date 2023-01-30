// Aula 01 - Manipulação de arquivos
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "./01_manipulacao_arquivos/arquivo.txt"
	// exemplo de criacao de arquivos

	// criando o arquivo pra escrita com a funcao Create do pacote os
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// exemplo de escrita de strings em um arquivo com a função WriteString
	// tamanho, err := f.WriteString("Hello World!")

	//  exemplo de escrita de bytes em um arquivo com a função Write
	tamanho, err := f.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %v bytes\n", tamanho)
	f.Close()

	// exemplo de leitura de arquivo

	// abrindo arquivo para leitura com a funcao Open do pacote os
	// arquivo, err := os.Open(filename)

	// abrindo e retornando os dados do arquiv com a funcao ReadFile do pacote os
	arquivo, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// convertedo para string porque o ReadFile vai retornar um array de bytes que representam o conteudo do arquivo
	fmt.Println("lendo conteudo do arquivo:", string(arquivo))

	// exemplo de leitura de um arquivo por streaming (ou seja, lendo aos poucos o conteudo do arquivo)
	// abordagem usada para arquivos muito grandes, para nao carregar tudo de uma vez na memoria e sobrecarregar a aplicacao

	arquivo2, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// le o conteudo do arquivo de forma bufferizada, ou seja, em pequenos pedacos por vez
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3) // cria um buffer pra ler 3 bytes por vez

	for {
		// retornar a possicao para ser lida como indice no buffer mais tarde
		indice, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println("lendo conteudo do arquivo por streaming:", string(buffer[:indice]))
	}

	// exemplo de remoção de arquivo com a funcao Remove do pacote os
	fmt.Println("removendo arquivo...")
	if err := os.Remove(filename); err != nil {
		panic(err)
	}
}
