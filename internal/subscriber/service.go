package subscriber

import (
	"analytic-reporting/internal/subscriber/dbhandler"
	"analytic-reporting/internal/subscriber/kafka"
	"analytic-reporting/internal/subscriber/realtimeorder"
	"fmt"
)

func InitSubscriberService() {
	chConn, err := dbhandler.ConnectClickHouse()
	if err != nil {
		fmt.Println("Error connecting to ClickHouse: %v\n", err)
	}

	master, err := kafka.ConnectKafkaConsumer()
	if err != nil {
		panic(err)
	}

	realtimeorder.InitRealtimeOrderModule(master, chConn)
}
