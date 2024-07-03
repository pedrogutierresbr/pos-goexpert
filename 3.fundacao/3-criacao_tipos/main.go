// criação de tipos
package main

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
	println(f)
}