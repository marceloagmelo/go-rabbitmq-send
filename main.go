package main

import (
	"log"

	"github.com/marceloagmelo/go-rabbitmq-send/app"
	"github.com/marceloagmelo/go-rabbitmq-send/config"
)

func main() {
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	log.Println("Servico escutando a 8080...")
	app.Run(":8080")

}
