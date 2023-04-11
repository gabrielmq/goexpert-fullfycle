// 02 - Trabalhando com Wait Groups
package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is runing\n", i, name)
		time.Sleep(time.Second)

		// indica para o wait group que o processo foi finalizado
		// e diminui a quantidade inicial das operacoes até que chega a zero
		// indicando que todas as tarefas já foram executadas
		wg.Done()
	}
}

func main() {
	// Wait Groups é uma forma de sincronizar a execução das tarefas/operações (threads)
	wg := sync.WaitGroup{}

	// adicionando a quantidade total de operações que serão executadas para o wait group saber
	wg.Add(25)

	// é necessário passar o waitgroup como um poiteiro, pois o Done() vai
	// diminindo o valor inicial de Add(), conforme as operações forem finalizadas,
	// se não passar por ponteiro o valor inicial de Add() nunca será decrementado
	// pois nao vamos estar trabalhando com a referencia da variavel na memoria
	// e sim com uma cópia
	go task("A", &wg)
	go task("B", &wg)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is runing\n", i, "anonima")
			time.Sleep(time.Second)
			wg.Done() // indica para o wait group que o processo da thread foi finalizado
		}
	}()

	// faz o wait group esperar até que todas as threads completem sua execução dentro do wait group
	// o Wait() identifica que não precisa esperar mais até que todos os processo infomados no Add()
	// cheguem a zero
	wg.Wait()
}
