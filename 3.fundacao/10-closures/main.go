// closures --> são as funções anônimas no go
package main

import "fmt"

func main() {
	total := func () int {
		return sum(1,3, 523,125,598,7458,562,35,954,6) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}