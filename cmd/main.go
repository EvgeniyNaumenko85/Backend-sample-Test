package main

import (
	"BST"
	"BST/configs"
	"BST/logger"
	"BST/pkg/handler"
	"BST/pkg/repository"
	"BST/pkg/service"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

// @title => BST API <=
// @version 1.0
// @description API server for Backend sample test Application

// @host localhost:9090
// @BasePath /

// @SecurityDefinitions BearerAuth
// @in header
// @name Authorization

func main() {
	configs.PutAdditionalSettings()
	logger.Init()
	logger.Info.Println("BST started")

	if err := initConfig(); err != nil {
		logger.Error.Println(err)
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Error.Println(err)
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	cfg1 := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.Username, config.Password, config.DBName, config.Port, config.SSLMode)

	repos := repository.NewRepository(cfg1)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(BST.Server)

	if err := srv.Run(ctx, viper.GetString("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}

	logger.Error.Println("BST exited")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
