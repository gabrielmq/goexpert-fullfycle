// Aula 01 - Primeiros passos

// arquivos .go por padrão devem ter uma declaração de package logo no inicio
// esse package deve ter o nome do diretório em que o arquivo esta,
// com exeção do package main, pois esse é o package principal que contem a funcao main que é o start da app
// todos os aquivos .go dentro de um diretório devem ter o mesmo nome de package, ou seja, o nome do diretório
package main

const helloWorld = "Hello World!"

// função principal dos arquivos .go, ela que diz para o go iniciar a aplicação
func main() {
	println(helloWorld)
}
