package tax

import "testing"

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

// testing.F é utilizado para testes de mutação
// Fuzz é um tipo de teste que vai testar a variação dos parametros em um função
func FuzzCalculateTax(f *testing.F) {
	// seed são valores passados para o fuzz entender o tipo de parametro que a função recebe para ele poder enviar os valores.
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1500.0}

	for _, amount := range seed {
		// podem ser informados mais de um parametro para o fuzz
		// f.Add(param1, param2)
		f.Add(amount) // dando exemplos de parametros para o fuzz
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		actualTax := CalculateTax(amount)
		if amount <= 0 && actualTax != 0 {
			t.Errorf("Received %f but but expected 0", actualTax)
		}

		if amount > 20000.0 && actualTax != 20 {
			t.Errorf("Received %f but but expected 20", actualTax)
		}
	})
}
