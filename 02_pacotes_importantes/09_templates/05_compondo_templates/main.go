// 09_5 - Compondo templates
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
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ParseFiles é uma funcao variadica e os ... indicam os parametros
		// template.New() usa um arquivo como base para depois parsear os demais
		t := template.Must(template.New("content.html").ParseFiles(templates...))

		cursos := Cursos{
			{"Go", 40},
			{"Java", 60},
			{"Python", 20},
		}

		if err := t.Execute(w, cursos); err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8282", nil))
}
