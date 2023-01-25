// definicao do pacote matematica por causa do diretorio matematica
package matematica

// funcoes iniciadas com letra maiuscula, são publicas e visiveis por quem importar o pacote matematica
func Soma[T int | float64](a, b T) T {
	return a + b
}

// funcoes iniciadas com letra minuscula, são privadas e visiveis apenas internamente no mesmo pacote/arquivo
func multiplica[T int | float64](a, b T) T {
	return a * b
}
