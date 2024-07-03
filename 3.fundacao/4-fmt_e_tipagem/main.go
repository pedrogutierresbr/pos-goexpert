// importando fmt e tipagem
package main

import "fmt"

const a = "Hello, world!"

type ID int // podemos criar tipos

var (
	b bool    = true
	c int     = 10
	d string  = "Pedro"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", f) // dessa forma imprime o tipo
	fmt.Printf("O valor de E é %v", f) // dessa forma imprime o valor
}