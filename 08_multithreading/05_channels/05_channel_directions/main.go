// 05 - Channel directions
// Uma boa prática é sempre definir a direção dos channels que são passados por parametro
// para garantir consistencia e ficar mais claro se o channel só vai receber ou enviar informaçoes
package main

import "fmt"

// Channel directions é o conceito de leitura ou envio de informações em um channel
// dependendo da direção da seta no channel, vamos estar lendo ou enviando infos para o channel
func main() {
	ch := make(chan string)
	go recebe("Hello", ch)
	ler(ch)
}

// <-chan send-only
// o parametro <-chan string, indica que o channel só vai enviar informação
func ler(dado <-chan string) {
	// <- a esquerda, indica que o channel estara entregando alguma info
	fmt.Println(<-dado)
}

// receive only chan<-
// o parametro ch chan<- string, indica que o channel só vai receber informação
func recebe(nome string, ch chan<- string) {
	// <- a direita, indica que o channel esta recebendo valor
	ch <- nome
}
