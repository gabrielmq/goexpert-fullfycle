package tax

import "testing"

// Teste de Benchmark
// testing.B é usado para testes de benchmark
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
