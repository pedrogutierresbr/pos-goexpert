// HttpClient com timeout
// crucial para seu sistema ser performático, é você estabelecer limites nas chamadas externas que seu sistema vai realizar
// exemplo, quanto tempo ela vai ficar esperando uma response de um serviço

// cuidado que esse código esta sendo reconhecido como trojan pelo Kaspersky
package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	resp, err := c.Get("https://google.com/")
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