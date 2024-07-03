// Defer --> serve para atrasar algum comportamento, ele ser feito por ultimo (segura para ser executado por ultimo)
package main

import (
	"io"
	"net/http"
)

func main(){
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close() // vai ser executado por ultimo

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))
}