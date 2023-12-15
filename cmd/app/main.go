package main

import (
	todo "github.com/kharitmaboy/simple-todo-list"
	"github.com/kharitmaboy/simple-todo-list/pkg/handler"
	"github.com/kharitmaboy/simple-todo-list/pkg/repository"
	"github.com/kharitmaboy/simple-todo-list/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}