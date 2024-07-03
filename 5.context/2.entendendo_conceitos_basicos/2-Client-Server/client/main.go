package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// vou realizar a requisição e se ela passar de 5 segundo, vai cancelar meu contexto, assim cancelando minha operação
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "Get", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// caso eu receba uma resposata, vamos imprimir no comand line Stdout
	io.Copy(os.Stdout, res.Body)
}

// consigo controlar o contexto tanto do lado do server quanto do lado do client

// Esse exemplo mostra o seguinte:
// fiz uma request e ai tinha 3 segundo de espera para o processamento. Como atingiu esse limite
// meu client desconetou do servidor. Servidor caiu o case Done() nesse caso, fazendo parar o processamento em andamento
// pois esse processamento não precisava mais cancelar, liberando assim a aplicação (server) para outros processamentos

// aqui esse exemplo foi pra requisição http, mas também pode ser usado para banco de dados