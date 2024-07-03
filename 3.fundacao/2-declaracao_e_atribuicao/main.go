// declaração e atribuição
package main

const a = "Hello, world!"

// consta não é possível alterar o valor

var (
	b bool    = true
	c int     = 10
	d string  = "Pedro"
	e float64 = 1.2
)

// var é possível alterar o valor

func main() {
	var a string = "X"
	f := "Y" // shorthand para criar, declarar e atribuir uma variável (ele faz a inferência), mas é quando pe feito pela primeira vez
	f = "ZZ" // para atribuir, você n]ao usa mais os :
	println(a)
	println(f)
}