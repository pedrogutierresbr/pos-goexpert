package main

import (
	"github.com/google/uuid"
	"github.com/pedrogutierresbr/pos-goexpert/7.packaging/3-trabalhando_com_go_mod_replace/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
