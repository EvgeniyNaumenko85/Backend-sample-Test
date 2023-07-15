package main

import (
	"BST"
	"BST/configs"
	"BST/logger"
	"BST/pkg/handler"
	"BST/pkg/repository"
	"BST/pkg/service"
	"context"
	"log"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := "host=localhost user=postgres password=Fe4ZOjoj dbname=postgres port=5432 sslmode=disable"
	repos := repository.NewRepository(cfg)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(BST.Server)

	if err := srv.Run(ctx, "9090", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}

	logger.Error.Println("BST exited")
}
