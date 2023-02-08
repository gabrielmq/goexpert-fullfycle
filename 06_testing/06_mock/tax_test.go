package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// o nome do teste deve sempre iniciar com Test
func TestCalculateTax(t *testing.T) {
	amount := 1000.0
	expectedTax := 10.0

	actualTax, err := CalculateTax(amount)

	assert.Nil(t, err)
	assert.Equal(t, expectedTax, actualTax)

	actualTax, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Equal(t, 0.0, actualTax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}

	// mockando a chamada da função SaveTax que quando passado determinado parametro
	// a funcao tera um retorno especifico
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	// informando ao mock que essa funcao pode ser chamada apenas 1 vez
	// repository.On("SaveTax", 10.0).Return(nil).Once()

	// informando ao mock que essa funcao pode receber qualquer valor
	// repository.On("SaveTax", mock.Anything()).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	// vai validar o assert
	repository.AssertExpectations(t)

	// validando a quantidade de vezes que a funcao SaveTax foi executada internamente
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
