package main

import (
	"fmt"
	"log"

	"github.com/aburizalpurnama/my-go-fiber/config"
	"github.com/aburizalpurnama/my-go-fiber/handler"
	"github.com/aburizalpurnama/my-go-fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	var conf config.MainConfig
	getEnv(&conf)

	handler := handler.NewHandler(conf)

	app := fiber.New(fiber.Config{
		ReadTimeout:  conf.Server.Timeout,
		WriteTimeout: conf.Server.Timeout,
		AppName:      "my go-fiber",
	})

	router.SetupRoutes(app, handler)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)))
}

func getEnv(config *config.MainConfig) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Failed to load config, unable to decode into struct, %v", err)
	}
}
