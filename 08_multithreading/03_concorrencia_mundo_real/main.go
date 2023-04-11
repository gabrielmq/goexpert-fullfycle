// 03 - Simulando concorrencia do mundo real
package main

import (
	"fmt"
	"net/http"
)

// variavel global, que pode sofre atualizações inesperadas
// em um ambiente concorrente por causa de race conditions entre as threads
var number int = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte(fmt.Sprintf("Voce e o visitando numero %d", number)))
	})

	// o servidor vai tratar um request por thread, nesse cenário podemos cair em race conditions
	// por causa da variavel global number, que é compartilhada entre todas as threads abertas
	http.ListenAndServe(":8000", nil)
}
