// Trabalhando com HTTP usando Contextos
// pacote context permite que a gente passe as informações dele, para diversas chamadas do nosso sistema e nos da a opção desses contextos sejam cancelados
// qundo contexto é cancelado, a gente para a operação na hora, para que a gente não fique gastando muito tempo

// contexto permite a gente conseguir pedir para nossos programas pararem a execuçao de algo que eles estejam fazendo, que eles não vão mais precisar fazer
// de acordo com o contexto que a aplicação esta rodando
package main

import (
	"context"
	"io"
	"net/http"
	"time"
)


func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp,err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}