// percorrendo arrays
package main

import "fmt"

const a = "Hello, world!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Pedro"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	// Array trabalha como variável, tem tamanho fixo, e você pode percorre-lo
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 20

	// i para índice
	// v para valor
	// range {array} para percorrer os itens do array
	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}
}

// %v para string
// %d para dígito