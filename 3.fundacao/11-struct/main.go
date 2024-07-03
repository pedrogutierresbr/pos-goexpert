// Struct --> para meios de comparação, seria o que mais se assemelha a uma classe no go (mas mainda sim, não é uma classe)
package main

import "fmt"


type Cliente struct  {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	pedro := Cliente{
		Nome: "Pedro",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t \n", pedro.Nome, pedro.Idade, pedro.Ativo)

	pedro.Ativo = false

	fmt.Println(pedro.Ativo)
}
