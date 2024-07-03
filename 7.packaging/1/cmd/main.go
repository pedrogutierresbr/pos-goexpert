package main

import (
	"fmt"

	"github.com/pedrogutierresbr/pos-goexpert/7.packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m)
	fmt.Println(m.Add())
	// fmt.Println(math.X)

	fmt.Println(m.C)
	m.C = 7
	fmt.Println(m.C)
}

// letra maiúscula está exportando
// letra minúscula não está exportando
// essa aula é pra exemplicar como usar os modificadores de acesso
// private, public ...
