// Customizando Request
// client http é diferente de objeto de request

// Bem comum de utilizar, principalmente quando vc quer customizar bastante a request, quando quer passar um baerer token, campos que quer trabalhar ...
package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req) // essa linha que é responsável pelo client pegar o objeto de request e executar a chamada
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