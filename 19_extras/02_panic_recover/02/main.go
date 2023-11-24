package main

import (
	"log"
	"net/http"
)

func recoverMiddlaware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic: %v\n", r)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("ERROR!")
	})

	log.Println("listen on port :3000")
	if err := http.ListenAndServe(":3000", recoverMiddlaware(mux)); err != nil {
		log.Fatalf("could not listen on :3000 %v", err)
	}
}
