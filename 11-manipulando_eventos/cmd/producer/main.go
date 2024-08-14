package main

import "github.com/pedrogutierresbr/pos-goexpert/11-manipulando_eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}

// para funcionar:

// crie uma fila chamada "orders" no rabbitmq adm
// entre em exchange
// entre na exchange padrão do rabbitmq: amq.direct
// dê um bind do queue orders na exchange amq.direct
