package app

import (
	"analytic-reporting/internal/subscriber/realtimeorder/adapters/clickhouse"
	kafkaadapter "analytic-reporting/internal/subscriber/realtimeorder/adapters/kafka"
	"analytic-reporting/internal/subscriber/realtimeorder/domain"
	"analytic-reporting/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RealtimeOrderService interface {
	CreateOrderRecord(message string)
}

type realtimeOrderService struct {
	repository       clickhouse.RealtimeOrderRepository
	realtimeConsumer kafkaadapter.RealTimeConsumer
}

func InitRealtimeOrderService(
	repository clickhouse.RealtimeOrderRepository,
	realtimeConsumer kafkaadapter.RealTimeConsumer,
) RealtimeOrderService {
	return &realtimeOrderService{
		repository:       repository,
		realtimeConsumer: realtimeConsumer,
	}
}

func (s *realtimeOrderService) CreateOrderRecord(message string) {
	parsedMessage := s.ProcessMessageString(message)
	orderRecord := domain.OrderRecord{
		Id:          uuid.New(),
		SaleId:      parsedMessage.ID,
		Quantity:    parsedMessage.Qty,
		SaleAmount:  parsedMessage.SaleAmount,
		Discount:    parsedMessage.Discount,
		Profit:      parsedMessage.Profit,
		ProfitRatio: parsedMessage.ProfitRatio,
		OrderId:     parsedMessage.OrderID,
		OrderDate:   parsedMessage.OrderDate,
		LocationId:  parsedMessage.LocationID,
		ProductId:   parsedMessage.ProductID,
		SegmentId:   parsedMessage.SegmentID,
		ProductName: parsedMessage.ProductName,
		SegmentName: parsedMessage.SegmentName,
		CreatedAt:   time.Now(),
	}

	err := s.repository.CreateOrderRecord(orderRecord)
	if err != nil {
		fmt.Println("Error while creating record", err)
	}
}

func (s *realtimeOrderService) ProcessMessageString(message string) model.SaleWithDetail {
	var saleData model.SaleWithDetail

	err := json.Unmarshal([]byte(message), &saleData)
	if err != nil {
		fmt.Println("error:", err)
	}

	return saleData
}
