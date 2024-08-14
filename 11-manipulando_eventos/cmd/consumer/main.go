package main

import (
	"fmt"

	"github.com/pedrogutierresbr/pos-goexpert/11-manipulando_eventos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}

// sempre levantar o container do rabbitmq antes
// criar uma fila no rabbitmq adm

// depois rodar a aplicação
