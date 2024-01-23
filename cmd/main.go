package main

import (
	"github.com/LLlE0/URL_shortener/handler"
	serv "github.com/LLlE0/URL_shortener/service"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	server := serv.NewServer()
	store := serv.NewStore()
	log.Print("The server started successfully at: " + viper.GetString("ip") + ":" + viper.GetString("port"))
	server.Run(viper.GetString("port"), viper.GetString("ip"), handler.InitRoutes(store))

}

// handle config file
func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
