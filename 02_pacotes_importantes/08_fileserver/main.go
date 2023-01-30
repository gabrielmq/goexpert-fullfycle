// 08 - FileServer
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// cria um servidor de arquivos estaticos informando o diretorio
	// onde os arquivos estao armazenados
	fileServer := http.FileServer(http.Dir("./public"))

	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Blog!"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
