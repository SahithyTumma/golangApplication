package pulsar

import (
	"context"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

var client pulsar.Client
var producer pulsar.Producer
var consumer pulsar.Consumer

func InitializePulsarClient() {
	var err error

	client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://10.96.73.196:6650",
	})
	if err != nil {
		log.Fatalf("Failed to create Pulsar client: %v", err)
	}

	producer, err = client.CreateProducer(pulsar.ProducerOptions{
		Topic: "topic",
	})
	if err != nil {
		log.Fatalf("Failed to create Pulsar producer: %v", err)
	}

	consumer, err = client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic",
		SubscriptionName: "Tsubscription",
	})
	if err != nil {
		log.Fatalf("Failed to create Pulsar consumer: %v", err)
	}
}

func ClosePulsarClient() {
	if consumer != nil {
		consumer.Close()
	}
	if producer != nil {
		producer.Close()
	}
	if client != nil {
		client.Close()
	}
}

func PublishMessage(message string) error {
	_, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(message),
	})
	return err
}

func ConsumeMessages() {
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}
		log.Printf("Received message: %s", string(msg.Payload()))
		consumer.Ack(msg)
	}
}
