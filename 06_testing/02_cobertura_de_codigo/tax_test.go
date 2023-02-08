package tax

import "testing"

// testing.T é usado para testes de unidade
// o nome do teste deve sempre iniciar com Test
func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectedTax := 5.0

	// o assert default é com if
	actualTax := CalculateTax(amount)

	if actualTax != expectedTax {
		t.Errorf("Expected %f but got %f", expectedTax, actualTax)
	}
}

// Exemplo de um teste em "batch" testando mais de um dado para a mesma funcao
// evitando de criar varios testes
func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expectedTax float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		actualTax := CalculateTax(item.amount)

		if actualTax != item.expectedTax {
			t.Errorf("Expected %f but got %f", item.expectedTax, actualTax)
		}
	}
}
