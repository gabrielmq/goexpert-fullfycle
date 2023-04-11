// 03 - Mutex e operações atomicas
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var number int64 = 0

func main() {
	// criando um Mutex para travar um processo para evitar rece conditions entre as threads
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// trava o processo para que apenas uma thread atualize o dado
		m.Lock()

		number++
		// outra forma de evitar race conditions
		// pode de baixo dos panos, esse atomic.AddInt64 já faz o lock/unlock
		// atomic.AddInt64(&number, 1)

		// libera o processo para que outra thread possa atualizar o dado
		m.Unlock()
		w.Write([]byte(fmt.Sprintf("Voce e o visitando numero %d", number)))
	})
	http.ListenAndServe(":8000", nil)
}
