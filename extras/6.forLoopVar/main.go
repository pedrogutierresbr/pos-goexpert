package main

import (
	"fmt"
)

func main() {
	//  agora é possível colocar o range direto, antes só era possível decalrando uma var e atribuindo valor nela
	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for range values {
		<-done
	}
}

// Aula feito para exemplificar a correção que fizeram na versão 1.22 no loop
// antes tinha um bug relacionado as variáveis, que ele repetia e exibia apenas as últimas
// devido a atribuição, ai tinha que colocar uma "gambiarra" para acertar o erro, ex:

// for _, v := range values {
// 	go func() {
// 		v := v
// 		fmt.Println(v)
// 		done <- true
// 	}()
// }

// ai dava tudo certo, agora com a nova versão corrigiram, e não precisa mais do v := v
