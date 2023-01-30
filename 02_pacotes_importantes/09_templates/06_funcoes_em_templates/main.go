// 09_6 - Mapeando funções nos templates
package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
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
		t := template.New("content.html")

		// faz o mapeamento das funcoes que podem ficar disponiveis para serem usadas dentro do template
		// sem esse mapeamento nao é possivel executar funcoes dentro do template
		t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})

		t = template.Must(t.ParseFiles(templates...))

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
