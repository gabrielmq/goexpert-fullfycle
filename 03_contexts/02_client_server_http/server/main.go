// Aula 02_1 - Contexts com server HTTP
package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// pegando context que o request possui
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	//
	select {
	// vai processar a req em até 5s
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))
		return

	// finaliza o contexto quando a req é cancelada
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		return
	}
}
