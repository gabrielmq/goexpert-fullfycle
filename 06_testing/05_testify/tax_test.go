package tax

import (
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
