package kafkaadapter

import (
	"fmt"

	"github.com/IBM/sarama"
)

type RealTimeConsumer interface {
	ConsumeRealTimeData() (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError)
}

type realTimeConsumer struct {
	topic  string
	master sarama.Consumer
}

func InitRealtimeConsumer(topic string, master sarama.Consumer) RealTimeConsumer {
	return &realTimeConsumer{
		topic:  topic,
		master: master,
	}
}

func (r *realTimeConsumer) ConsumeRealTimeData() (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	messages := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)

	partitions, err := r.master.Partitions(r.topic)
	if err != nil {
		close(messages)
		close(errors)
		fmt.Println("Failed to create partition consumer:", err)
		return messages, errors
	}

	partitionConsumer, err := r.master.ConsumePartition(r.topic, partitions[0], sarama.OffsetOldest)
	if err != nil {
		close(messages)
		close(errors)
		fmt.Println("Failed to create partition consumer:", r.topic, err)
		return messages, errors
	}

	go func(partitionConsumer sarama.PartitionConsumer) {
		defer close(messages)
		defer close(errors)
		defer partitionConsumer.Close()

		for {
			select {
			case msg := <-partitionConsumer.Messages():
				messages <- msg
			case err := <-partitionConsumer.Errors():
				errors <- &sarama.ConsumerError{
					Topic:     r.topic,
					Partition: 0,
					Err:       err,
				}
			}
		}
	}(partitionConsumer)

	return messages, errors
}
