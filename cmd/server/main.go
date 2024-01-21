package main

import (
	"context"
	"log"

	"github.com/lidiagaldino/desafio-backend/cmd/server/config"
	"github.com/lidiagaldino/desafio-backend/internal/infra/database"
	"github.com/lidiagaldino/desafio-backend/internal/infra/web"
	"github.com/spf13/viper"
)

func loadConfiguration() (config config.Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	return
}

func main() { 
	config, err := loadConfiguration()
	if err != nil {
		log.Fatal(err)
  }
	client, err := database.Connect(config.DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	web.Initialize()
 }
