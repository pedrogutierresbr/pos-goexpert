// Aulas
// Fuzzing
package tax

import "testing"

func FuzzCalculateTax(f *testing.F) {
	// valores a serem dados como exemplo para o teste
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}

// Esse cara vc alimenta ele no seed, para ele ter uma base
// depois que ele pega essa base ele sai passando um monte de valor até tentar quebrar a app
// ai vc consegue ir corrigindo sua lógica

// Para testar:
// arquivo tem que ser ser: nome + _test
// nome do método a ser testado: Fuzz...

// Forma de rodar o teste --> go test -fuzz=.

// Forma de rodar o teste e garantir que seja apenas o fuxx --> go test -fuzz=. -run=^#

// Forma de rodar o teste e determinar um tempo max pra rodar --> go test -fuzz=. -fuzztime=5s -run=^#
