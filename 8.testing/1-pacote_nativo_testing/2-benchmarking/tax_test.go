// Aulas
// Trabalhando com benchmarking
package tax

import "testing"

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// Teste de benchmark
// arquivo tem que ser ser: nome + _test
// nome do método a ser testado: Benchmark...
// argumento sempre vai ser (b *testing.B)

// Forma mais simplificada (porém roda todos os testes do arquivo)
// --> go test -bench=.

// Forma que roda apenas testes de benchmark (usando regex)
// --> go test -bench=. -run=^#

// Forma que mostra linha a linha o teste sendo executado, limitando ao número que vc passar no count
// --> go test -bench=. -run=^# -count={numero}

// Forma roda apenas a quantidade de tempo estipulada, para cada teste
// --> go test -bench=. -run=^# -count={numero} benchtime=3s

// Forma de exibir dados sobre alocação de memória junto
// --> go test -bench=. -run=^# -benchmem

// Menu de testes (doc)
// --> go help test
