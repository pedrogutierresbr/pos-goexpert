// maps --> sãos que bem os objetos do js
package main

import (
	"fmt"
)

func main() {
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Nicole": 3589}
	delete(salarios, "Nicole")
	salarios["Nic"] = 5000

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1("Wesley") = 50012

	for nome, salario := range salarios{
		fmt.Printf("O salário de %s é %d\n",nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}