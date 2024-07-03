package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <- time.After(5 * time.Second):
		// imprime no comand line Stdout
		log.Println("Request processada com sucesso!")
		// imprime no browser
		w.Write([]byte("Request processada com sucesso!"))
	case <- ctx.Done(): // acontece caso a requisição seja cancelada pelo lado do client
		// imprime no comand line Stdout
		log.Println("Request cancelada pelo cliente!")
	}
}	

// Na prática evita que a gente fique fazendo processamentos desnecessários conforme o tempo vai passando