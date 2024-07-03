// File server --> criando um servidor de arquivos estáticos (imagens, html, css ...)
package main

import (
	"log"
	"net/http"
)

func main()  {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()

	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo do Blog!"))
	})

	log.Fatal((http.ListenAndServe(":8080", mux)))
}