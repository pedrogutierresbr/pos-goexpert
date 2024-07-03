// Ponteiros --> quando usar ponteiros
package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	var c = soma(&minhaVar1, &minhaVar2)
	println(c)
	println(minhaVar1)
	println(minhaVar2)
}


// Quando não usar ponteiros?
// quando você quer passar uma cópia dos dados, fazer alguma utilização

// Quando usar ponteiros?
// quando você quer trabalhar com valores mutáveis (quando o valor precisa mudar em diversos lugares)