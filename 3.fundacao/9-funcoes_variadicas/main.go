// funções variádicas --> quando não sabe quantos argumentos a função recebe
package main

import "fmt"

func main() {
	fmt.Println(sum(1,3, 523,125,598,7458,562,35,954,6))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}