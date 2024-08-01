package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

// Exemplo a ser feito, uma servidor web que esta recebendo muitas reqs de diversos users simultaneamente

// Primeiro ele roudou o server, depois em outro terminal:

// na aula foi feito um teste de benchmark, usando ApacheBench
// rodou 10000 req no curl http://localhost:3000/

// comando: ab -n 10000 -c 100 http://localhost:3000/   --> feito 10000 reqs por 100 pessoal simultâneas

// depois ele fez uma manual, e visualizou que não foi retornando 10001
// dai ele adicionou a linha time.Sleep(300 * time.Millisecond)

// então reiniciou o server e fez o teste de bench novamente
// depois, rodou uma vez manual e foi averiguado que não chegou nos 10000

// então constatado que tivemos problemas de concorrÊncia

// digamos, que esse exemplo, seja um exemplo de race condition
