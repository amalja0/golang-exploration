package realtimeorder

import (
	"analytic-reporting/internal/subscriber/realtimeorder/adapters/clickhouse"
	kafkaadapter "analytic-reporting/internal/subscriber/realtimeorder/adapters/kafka"
	"analytic-reporting/internal/subscriber/realtimeorder/app"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/IBM/sarama"
)

func InitRealtimeOrderModule(master sarama.Consumer, chConn driver.Conn) {
	topic := "sales"
	consumers := kafkaadapter.InitRealtimeConsumer(topic, master)

	sqlPath := "sql/realtime_order.sql"
	realtimeOrderRepository := clickhouse.InitRealtimeRepository(sqlPath, &chConn)
	realtimeOrderService := app.InitRealtimeOrderService(realtimeOrderRepository, consumers)

	err := realtimeOrderRepository.InitDB()
	if err != nil {
		fmt.Println("Error initializing DB: %v\n", err)
	}

	msgChan, _ := consumers.ConsumeRealTimeData()

	for msg := range msgChan {
		stringified := string(msg.Value)
		realtimeOrderService.CreateOrderRecord(stringified)
	}
}
