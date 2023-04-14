package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// Cria uma conexão com rabbit
func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// cria um channel com rabbitmq
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

// amqp.Delivery tem as informações da msg a ser consumida
func Consume(queue string, ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, exName string, msg string) error {
	return ch.Publish(
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}
