// 09_1 - Iniciando com templates
package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}

	// Criando um template dinamico para substituir valores
	tmp := template.New("CursoTemplate")

	// Criando a saida de texto no template
	tmp, err := tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
	// executando template para sair no terminal
	if err := tmp.Execute(os.Stdout, curso); err != nil {
		panic(err)
	}
}
