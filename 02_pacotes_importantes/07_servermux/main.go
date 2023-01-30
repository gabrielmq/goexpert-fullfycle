// 07 - ServerMux
package main

import "net/http"

type blog struct{}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Blog"))
}

func main() {
	// cria uma multiplexer para atachar varias rotas no servidor
	mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello ServerMux"))
	// })

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{})

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello ServerMux"))
}
