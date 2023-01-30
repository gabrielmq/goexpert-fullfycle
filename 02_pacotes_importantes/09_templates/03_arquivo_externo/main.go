// 09_3 - Arquivo externo
package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	// template.Must já encapsula o tratamento de erro
	// parseando os valores para html com ParseFiles
	t := template.Must(template.New("template.html").ParseFiles("./template.html"))

	// executando template indicando um output
	if err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 60},
		{"Python", 20},
	}); err != nil {
		panic(err)
	}
}
