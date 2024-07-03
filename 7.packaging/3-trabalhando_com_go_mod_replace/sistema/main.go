package main

import "github.com/pedrogutierresbr/pos-goexpert/7.packaging/3-trabalhando_com_go_mod_replace/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}

// comando para direcionar para pasta local
// go mod edit -replace {url}={url relativa do pacote (pasta)}
// depois tem que fazer um go mod tidy

//Usar apenas para ambiente local, pois pode dar pau, devido a não ter subido as dependencias
