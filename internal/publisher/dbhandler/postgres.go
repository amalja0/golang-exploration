package dbhandler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ConnectPostgres() (*sql.DB, error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/publisher/config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.name"),
		viper.GetString("db.ssl_mode"),
	)

	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	//defer db.Close()

	return db, err
}
