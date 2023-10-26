package main

import (
	"fmt"
	"log"

	"github.com/aburizalpurnama/my-go-fiber/config"
	"github.com/aburizalpurnama/my-go-fiber/handler"
	"github.com/aburizalpurnama/my-go-fiber/model/domain"
	"github.com/aburizalpurnama/my-go-fiber/repository"
	"github.com/aburizalpurnama/my-go-fiber/router"
	"github.com/aburizalpurnama/my-go-fiber/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var conf config.MainConfig
	getEnv(&conf)

	// tied up dependecies
	db := connectDB(conf.Database)
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)
	handler := handler.NewHandler(conf)

	app := fiber.New(fiber.Config{
		ReadTimeout:  conf.Server.Timeout,
		WriteTimeout: conf.Server.Timeout,
		AppName:      "my go-fiber",
	})
	router.SetupRoutes(app, handler, userHandler)

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

func connectDB(config config.DatabaseConfig) *gorm.DB {
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  config.Dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB, err:%s", err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("Failed to migrate, err:%s", err)
	}

	return db
}
