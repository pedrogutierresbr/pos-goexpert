// Interface
package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct  {
	Nome string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado! \n", c.Nome)
}

func Desativacao(pessoa Pessoa){
	pessoa.Desativar()
}

func main() {
	pedro := Cliente{
		Nome: "Pedro",
		Idade: 30,
		Ativo: true,
	}

	println(pedro)

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
}
