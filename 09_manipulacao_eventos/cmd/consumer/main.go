package main

import (
	"fmt"

	"github.com/gabrielmq/events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	queue := "minhafila"
	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(queue, ch, msgs)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		// tira a msgs da fila, dizendo que ela já foi processada
		msg.Ack(false)
	}
}
