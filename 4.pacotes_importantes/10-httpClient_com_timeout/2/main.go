// Trabalhando com Post
// buffer é um espaço cheio de informação
package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "pedro"}`)) // é necessário bufferizar para realizar a operação de POST
	resp, err := c.Post("https://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}