// Composição de Structs
package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct  {
	Nome string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	pedro := Cliente{
		Nome: "Pedro",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t \n", pedro.Nome, pedro.Idade, pedro.Ativo)

	pedro.Ativo = false
	pedro.Cidade = "Uberaba"
	pedro.Endereco.Estado = "MG"

	fmt.Println(pedro)
}
