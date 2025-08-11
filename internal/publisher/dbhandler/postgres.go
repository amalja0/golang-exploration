package dbhandler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ConnectPostgres() (*sql.DB, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("PG_USER"),
		viper.GetString("PG_PASSWORD"),
		viper.GetString("PG_DB_NAME"),
		viper.GetString("PG_SSL_MODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return db, err
}
