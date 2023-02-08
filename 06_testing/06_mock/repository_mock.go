package tax

import "github.com/stretchr/testify/mock" // pacote de mock do testify

// Mock que representa o Repositorio
type TaxRepositoryMock struct {
	mock.Mock // vai permitir simular os metodos do repositorio
}

func (t *TaxRepositoryMock) SaveTax(tax float64) error {
	args := t.Mock.Called(tax) // registra no mock que a função foi chamada
	return args.Error(0) // retorna error porque é o tipo de retorno
}
