package routers

import (
	"net/http"

	"github.com/marceloagmelo/go-rabbitmq-send/controllers"
)

//CarregaRotas  as rotas
func CarregaRotas() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/health", controllers.Health)
	http.HandleFunc("/v1/new", controllers.New)
	http.HandleFunc("/v1/insert", controllers.Insert)
}
