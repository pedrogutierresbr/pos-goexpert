// Ponteiros
package main

func main() {
	// Memória -> Endereço -> Valor
	// usar o &{variável} é descobrir o endereçamento de memória
	// forma de ir não pelo valor mas pelo endereço da varável

	a := 10
	// println(&a)
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(a)
}
