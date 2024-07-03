// Ponteiros e Structs
package main

import "fmt"

type Conta struct {
	saldo int
}

// aqui como tenho uma cópia, vai alterar ali dentro, porém fora da função, não terá alterção
func (c Conta) simular(valor int) int {
	c.saldo += valor
	fmt.Printf("Valor da simulação: %d \n", c.saldo)
	return c.saldo
}

// Se eu fizer da forma abaixo, usando o ponteiro, o meu valor armazenado em memória sera alterado
// diferente da forma acima que é feito uma cópia do valor, e refletida apenas naquele escopo local

// func (c *Conta) simular(valor int) int {
// 	c.saldo += valor
// 	fmt.Printf("Valor da simulação: %d \n", c.saldo)
// 	return c.saldo
// }

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	fmt.Println(conta.saldo)
}

