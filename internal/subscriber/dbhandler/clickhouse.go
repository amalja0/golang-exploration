package dbhandler

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/spf13/viper"
)

func ConnectClickHouse() (driver.Conn, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(
		"%s:%s",
		viper.GetString("CH_HOST"),
		viper.GetString("CH_PORT"),
	)
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{address},
		Auth: clickhouse.Auth{
			Database: viper.GetString("CH_DB_NAME"),
			Username: viper.GetString("CH_USER"),
			Password: viper.GetString("CH_PASSWORD"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx := context.Background()
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		log.Fatal(err)
	}

	fmt.Println("Connected to ClickHouse!")

	return conn, err
}
