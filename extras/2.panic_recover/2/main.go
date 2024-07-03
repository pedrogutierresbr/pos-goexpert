// Exemplo 2: webserver simples
package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

func recoveredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic: %v\n", r)
				debug.PrintStack() // utilize isso se vc quiser exibir o stack trace
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	if err := http.ListenAndServe(":3000", recoveredMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen on %s: %v \n", ":3000", err)
	}
}

// use o curl para acessar os endpoints
