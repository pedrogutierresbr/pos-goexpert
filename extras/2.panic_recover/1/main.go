// Panic recover --> um pânico acontece, a gente pega esse pânico, vê qual pânico aconteceu
// e ai tomamos alguma ação
package main

import "fmt"

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

func main() {
	defer func() {
		if r := recover(); r != nil { // vai receber o pânico, e baseado nele, vai tomar uma ação
			if r == "panic1" {
				fmt.Println("panic1 recovered")
			}
			if r == "panic2" {
				fmt.Println("panic2 recovered")
			}
		}
	}()

	// panic1() //faça o chaveamento para rodar o exemplo
	panic2()
}
