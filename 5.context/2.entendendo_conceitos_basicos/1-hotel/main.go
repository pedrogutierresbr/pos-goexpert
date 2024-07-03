// Context -- entendendo conceitos básicos
// Exemplo booking de um hotel

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() // iniciei um context em branco
	ctx, cancel := context.WithTimeout(ctx, time.Second*3) // peguei o context e adicionei ele como um timeout de 3 segundos 
	defer cancel() // defer para cancelar por ultimo

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): // daria um Done quando passar meu timeout de 3 segundos, se passar esse tempo ele vai falar que o tempo foi esgotado e ele cancelou a reserva
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <- time.After(2 * time.Second): // se o tempo der 5 segundo e ele não for cancelado, vai falar que a reserva foi feita (poderia ser qualquer outra operação)
		fmt.Println("Hotel booked.")
		return
	}
}

// WithTimeout --> se o context expira, ele é cancelado, se ele é cancelado recebemos um Done, que executa o .Done()