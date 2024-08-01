package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

// Exemplo a ser feito, uma servidor web que esta recebendo muitas reqs de diversos users simultaneamente (aula anterior)

// exemplificou o uso do Mutex, mencionado nas aulas de "slide", para solucionar o race condition

// ele repetiu o experimento, da aula passada
// rodou o teste de benchmark e depois fez uma consulta normal
// ai apareceu o 10001

// O go possui um recurso, usado no modo de desenvolvimento, para tentar averiguar s epossui race condition
// go run -race nome_programa.go

// Ele rodou esse comando, depois rodou o teste de benchmark em outro terminal
// ai no terminal principal, estourou um erro de DATA RACE

// ======================================================================================================================

// depois ele comentou a parte de mutex e apresentou um pacote chamado atomic
// ai ele fez algo, que ele chamou de soma atomica
// rodou o servidor, em outro terminal ele rodou o teste de benchmark
// depois fez uma consulta solo e teve o resultado de 10001
