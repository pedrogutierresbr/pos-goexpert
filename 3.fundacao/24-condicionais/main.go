// Condicionais
package main

func main() {

	// No golang não existe else if

	a := 1
	// b := 2
	// c := 5

	// if a > b {
	// 	println(a)
	// } else {
	// 	println(b)
	// }

	// if a > b && c > a {
	// 	println("a > b && c > a")
	// } else {
	// 	print("Inválido")
	// }

	// if a > b || c > a {
	// 	println("caiu no or")
	// } else {
	// 	print("Inválido")
	// }

	// Switch vc faz em cima de uma variável
	switch a {
	case 1:
		println("a")
	case 2: 
		println("b")
	case 3:
		println("c")
	default:
		println("default")
	}
}