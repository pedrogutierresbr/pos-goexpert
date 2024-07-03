// Pacotes
package main

import (
	"aula-modulos/matematica"
	"fmt"
)


func main() {
	s := matematica.Soma(10,20)
	carro := matematica.Carro{Marca: "Chevrolet"}

	fmt.Printf("Resultado %v\n", s)
	fmt.Println(matematica.A)
	fmt.Println(carro)
	fmt.Println(carro.Andar())
}

