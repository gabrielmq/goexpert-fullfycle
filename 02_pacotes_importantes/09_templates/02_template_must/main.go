// 09_2 - Template.Must
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

	// template.Must já encapsula o tratamento de erro
	// e o parse já é feito de uma só vez
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))

	// executando template indicando um output
	if err := t.Execute(os.Stdout, curso); err != nil {
		panic(err)
	}
}
