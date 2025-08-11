package app

import (
	postgresadapter "analytic-reporting/internal/publisher/sale/adapters/postgres"
	"analytic-reporting/internal/publisher/sale/domain"
	"analytic-reporting/model"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type Service interface {
	CreateSales(sale *domain.CreateSales) (int32, error)
	GetSales(queryParams postgresadapter.QueryParams) ([]model.SaleWithDetail, error)
}

type service struct {
	repository    postgresadapter.Repository
	kafkaProducer sarama.SyncProducer
}

func NewService(r postgresadapter.Repository, kafkaProducer sarama.SyncProducer) Service {
	return &service{
		repository:    r,
		kafkaProducer: kafkaProducer,
	}
}

func (s *service) CreateSales(sale *domain.CreateSales) (int32, error) {
	return s.repository.CreateSales(sale)
}

func (s *service) GetSales(queryParams postgresadapter.QueryParams) ([]model.SaleWithDetail, error) {
	sales, err := s.repository.GetSales(queryParams)
	for _, sale := range sales {
		bytes, err := json.Marshal(sale)
		if err != nil {
			fmt.Println(err)
		}
		s.PushSaleDataToQueue("sales", bytes)
	}
	return sales, err
}

func (s *service) PushSaleDataToQueue(topic string, message []byte) error {
	producer := s.kafkaProducer
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
