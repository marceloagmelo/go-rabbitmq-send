package controllers

import (
	"html/template"
	"log"
	"net/http"

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
