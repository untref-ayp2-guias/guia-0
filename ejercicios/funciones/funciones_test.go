package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImprimirPolinomio(t *testing.T) {
	coeficientes := []float64{3.0, 2.0, 1.0}
	resultadoEsperado := "3.0 + 2.0 x + 1.0 x**2"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(t, resultadoEsperado, resultado)
}
func TestImprimirPolinomioVacio(t *testing.T) {
	coeficientes := []float64{}
	resultadoEsperado := "El array está vacío"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(t, resultadoEsperado, resultado)
}
