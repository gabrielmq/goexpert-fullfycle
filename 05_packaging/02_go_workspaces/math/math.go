package math

// Struct publica por iniciar com letra maiuscula
type Math struct {
	a int
	b int
}

func NewMath(a, b int) *Math {
	return &Math{a, b}
}

// funcao publica por inicar com letra maiuscula
func (m *Math) Add() int {
	return m.a + m.b
}
