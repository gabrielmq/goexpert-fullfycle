// Aula 21 - Pacotes e modulos
package main

// pacotes customizados só serao importados se o projeto estiver usando modulos, cas contraio o go entende que o pacote vai estar dentro de go/src
// se o projeto estiver dentro da pasta go/src nao é necessario uar modulos
import (
	"fmt"

	"goexpert/matematica" // os imports de pacotes customizados, devem iniciar com o nome do modulo/diretorio
)

// também se aplica a variaveis e structs
// variaveis iniciadas com letra maiuscula, são publicas e visiveis por quem importar o pacote matematica
var A string = "a"

// Struct publica para fora do pacote
type Carro struct{}

// Struct privada e visivel apenas dentro do pacote
type teste struct{}

func main() {
	// usano a funcao Soma do pacote matematica
	soma := matematica.Soma(2, 2)
	// mult := matematica.multiplica(2, 2) nao funciona porque a funcao multiplica nao é privada
	fmt.Printf("Resultado: %v", soma)
}
