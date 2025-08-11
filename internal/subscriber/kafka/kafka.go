package kafka

import (
	"fmt"
	"strings"

	"github.com/IBM/sarama"
)

func ConnectKafkaConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.ClientID = "kafka-consumer-client"
	config.Consumer.Return.Errors = true

	brokers := []string{"localhost:19092"}

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		fmt.Println("Error connecting to kafka")
		panic(err)
	}

	fmt.Println("Connected to kafka")

	return master, err
}

func consume(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
		if nil != err {
			fmt.Printf("Topic %v Partitions: %v", topic, partitions)
			panic(err)
		}
		fmt.Println(" Start consuming topic ", topic)
		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError
					fmt.Println("consumerError: ", consumerError.Err)

				case msg := <-consumer.Messages():
					consumers <- msg
					fmt.Println("Got message on topic ", topic, msg.Value)
				}
			}
		}(topic, consumer)
	}

	return consumers, errors
}
