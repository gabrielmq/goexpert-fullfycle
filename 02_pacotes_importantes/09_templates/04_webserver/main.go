// 09_4 - Templates com webserver
package main

import (
	"html/template"
	"log"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// template.Must já encapsula o tratamento de erro
		// parseando os valores para html com ParseFiles
		t := template.Must(template.New("template.html").ParseFiles("./template.html"))

		cursos := Cursos{
			{"Go", 40},
			{"Java", 60},
			{"Python", 20},
		}

		// executando template indicando um output
		if err := t.Execute(w, cursos); err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8282", nil))
}
