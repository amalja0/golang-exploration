package kafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

func ConnectKafkaProducer() (sarama.SyncProducer, error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/publisher/config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	brokerUrl := fmt.Sprintf(
		"%s:%s",
		viper.GetString("kafka.host"),
		viper.GetString("kafka.port"),
	)
	brokersUrl := []string{brokerUrl}
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
