package main

import (
	todo "github.com/kharitmaboy/simple-todo-list"
	"github.com/kharitmaboy/simple-todo-list/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
