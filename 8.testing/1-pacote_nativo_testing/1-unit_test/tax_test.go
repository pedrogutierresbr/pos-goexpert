// Aulas
// Iniciando com testes automatizados
// Testando em Batch (lote)
package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expect %f but got %f", item.expected, result)
		}
	}
}

// Para testar:
// arquivo tem que ser ser: nome + _test
// nome do método a ser testado: Test...
// argumento sempre vai ser (t *testing.T)

// Para testar mais de 1 valor:
// método a ser testado tem que começar com Test...Batch
// argumento sempre vai ser (t *testing.T)

// As estruturas a cima servem de exemplo para testar um valor ou um slice de valores

// Forma mais simplificada --> go test .

// Forma mais verbosa (gostei mais) --> go test -v

// Faz o teste e mostra a % de cobertura --> go test -coverprofile=coverage.out

// Exibe uma interface html, mostrando o código, destacando caso alguma parte não tenha sido coberta
// go tool cover -html=coverage.out
