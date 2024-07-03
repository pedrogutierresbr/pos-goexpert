// Compilando projetos
package main

func main() {

	// Use --> go build {nome do pacote} --> build pro SO vigente

	// GOOS=windows go build {nome do pacote}  --> builda para windows (.exe) 
	// GOOS=linux go build {nome do pacote}  --> builda para linux 
	// GOOS=darwin go build {nome do pacote}  --> builda para mac 

	// GOOS={sistema operacional} GOARCH={arquitetura do processador} go build {nome do pacote}

	// go tool dist list --> mostra as extensoes possiveis

	// go env GOOS GOARCH --> mostra as configs do seu pc vigente

	// sai o binário final, pronto pra usar

	a := 1

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