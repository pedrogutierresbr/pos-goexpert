// Para essa aula, precisa da versão 1.22, pois trabalha com atualizaçõe que vieram nela
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("Get /books/{id}", GetBookHandler) // pega o parametro declarado
	// mux.HandleFunc("Get /books/dir/{d...}", BooksPathHandler) // pega qualquer coisa que precede o d
	// mux.HandleFunc("Get /books/{$}", BooksHandler) // pega exatamente até o limite estabelecido pelo {$}, qualquer coisa depois do {$} ele não pega

	// mux.HandleFunc("Get /books/precedence/latest", BooksPrecedenceHandler) // mais específico que o abaixo
	// mux.HandleFunc("Get /books/precedence/{x}", BooksPrecedence2Handler) // menos específico que o acima

	// mux.HandleFunc("Get /books/{id}", BooksPrecedenceHandler) // entre essa e a debaixo, ele estoura um panic dizendo que tem conflito, o pattern encontrou conflito pois amboes tem o mesmo grau de especificidade
	// mux.HandleFunc("Get /{x}/latest", BooksPrecedence2Handler)

	http.ListenAndServe(":9000", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book " + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirpath := r.PathValue("d") // Access captured directory path segments as slice
	fmt.Fprintf(w, "Accessing directory path: %s\n", dirpath)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}

func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence"))
}

func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence 2"))
}
