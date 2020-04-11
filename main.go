package main

import (
	"log"
	"net/http"

	"github.com/marceloagmelo/go-rabbitmq-send/routers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	routers.CarregaRotas()

	log.Println("Servico escutando a 8080...")
	http.ListenAndServe(":8080", nil)
}
