package main

import "github.com/gabrielmq/events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	ex := "amq.direct"
	rabbitmq.Publish(ch, ex, "Hello Rabbit!")
}
