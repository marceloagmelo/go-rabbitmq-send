package controllers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/marceloagmelo/go-rabbitmq-send/lib"
	"github.com/marceloagmelo/go-rabbitmq-send/models"
)

var view = template.Must(template.ParseGlob("views/*.html"))

//Home é a página inicial da aplicação
func Home(w http.ResponseWriter, r *http.Request) {
	mensagemErro := ""

	var mensagens []models.Mensagem

	if err := models.MensagemModel.Find().All(&mensagens); err != nil {
		mensagemErro = err.Error()
	}

	data := map[string]interface{}{
		"titulo":    "Lista de Mensagens",
		"mensagens": mensagens,
		"mensagem":  mensagemErro,
	}

	view.ExecuteTemplate(w, "Index", data)
}

//New página de edição de uma nova mensagem
func New(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"titulo":   "Nova Mensagem",
		"mensagem": "",
	}

	view.ExecuteTemplate(w, "New", data)
}

//Insert mensagem
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		titulo := r.FormValue("titulo")
		texto := r.FormValue("texto")

		if titulo != "" && texto != "" {
			var interf models.Metodos

			var mensagemForm models.Mensagem
			mensagemForm.Titulo = titulo
			mensagemForm.Texto = texto
			mensagemForm.Status = 1

			interf = mensagemForm

			mensagem := interf.Criar()
			if mensagem == "" {
				log.Println(mensagem)
			}
		}

	}
	http.Redirect(w, r, "/", 301)
}

//Health testa conexão com o mysql e rabbitmq
func Health(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")

	var mensagens []models.Mensagem
	if err := models.MensagemModel.Find().All(&mensagens); err != nil {
		log.Fatalf("%s: %s", "Erro ao conectar com o banco de dados", err)
	}

	conn, _ := lib.ConectarRabbitMQ()
	defer conn.Close()

	data := map[string]interface{}{
		"titulo": "Lista de Mensagens",
		"hora":   hora,
	}

	view.ExecuteTemplate(w, "Health", data)
}
