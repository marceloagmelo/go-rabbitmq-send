package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/marceloagmelo/go-rabbitmq-send/lib"
	"github.com/marceloagmelo/go-rabbitmq-send/models"
	"upper.io/db.v3"
)

var view = template.Must(template.ParseGlob("views/*.html"))

//Home é a página inicial da aplicação
func Home(db db.Database, w http.ResponseWriter, r *http.Request) {

	var mensagens []models.Mensagem
	var mensagemModel = db.Collection("mensagem")

	if err := mensagemModel.Find().All(&mensagens); err != nil {
		log.Printf("Home(): %s: %s", "Recuperando as mensagens", err)
	}

	data := map[string]interface{}{
		"titulo":    "Lista de Mensagens",
		"mensagens": mensagens,
	}

	err := view.ExecuteTemplate(w, "Index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
func Insert(db db.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "form error", http.StatusInternalServerError)
			return
		}

		titulo := r.FormValue("titulo")
		texto := r.FormValue("texto")

		if titulo != "" && texto != "" {
			var mensagemModel = db.Collection("mensagem")
			var interf models.Metodos

			var mensagemForm models.Mensagem
			mensagemForm.Titulo = titulo
			mensagemForm.Texto = texto
			mensagemForm.Status = 1

			interf = mensagemForm

			//err := interf.Criar(mensagemModel)
			if err := interf.Criar(mensagemModel); err != nil {
				//http.Error(w, err.Error(), http.StatusInternalServerError)
				mensagemErro := fmt.Sprintf("Insert(): %s: %s", "Erro ao criar a mensagem", err)
				data := map[string]interface{}{
					"titulo":       "Lista de Mensagens",
					"mensagemErro": mensagemErro,
				}

				err := view.ExecuteTemplate(w, "Erro", data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}

		}
		log.Println("O envio da mensagem realizado com sucesso!")

		mensagem := fmt.Sprintf("Envio da mensagem realizado com sucesso!")
		data := map[string]interface{}{
			"titulo":   "Lista de Mensagens",
			"mensagem": mensagem,
		}

		err = view.ExecuteTemplate(w, "Sucesso", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

//Health testa conexão com o mysql e rabbitmq
func Health(db db.Database, w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")

	var mensagemModel = db.Collection("mensagem")
	var mensagens []models.Mensagem
	if err := mensagemModel.Find().All(&mensagens); err != nil {
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

//RestTodasMensagens gravadas
func RestTodasMensagens(db db.Database, w http.ResponseWriter, r *http.Request) {
	var mensagens []models.Mensagem
	var mensagemModel = db.Collection("mensagem")
	if err := mensagemModel.Find().All(&mensagens); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, mensagens)
}

//RestUmaMensagem recuperar mensagem
func RestUmaMensagem(db db.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var mensagem models.Mensagem
	var mensagemModel = db.Collection("mensagem")

	resultado := mensagemModel.Find("id=?", id)
	if count, err := resultado.Count(); count < 1 {
		log.Printf("UmaMensagem(): %s: %s", "Mensagem não recuperada", err)
		mensagem := fmt.Sprintf("UmaMensagem(): Mensagem [%v] não recuperada!", id)
		log.Printf(mensagem)

		respondError(w, http.StatusNotFound, mensagem)
		return
	}

	if err := resultado.One(&mensagem); err != nil {
		log.Printf("UmaMensagem(): %s: %s", "Mensagem não recuperada", err)
		mensagem := fmt.Sprintf("UmaMensagem(): Mensagem [%v] não recuperada!", id)
		log.Printf(mensagem)

		respondError(w, http.StatusNotFound, mensagem)
		return
	}

	respondJSON(w, http.StatusCreated, mensagem)
}

//RestEnviarMensagem recuperar mensagem
func RestEnviarMensagem(db db.Database, w http.ResponseWriter, r *http.Request) {
	var novaMensagem models.Mensagem

	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("EnviarMensagem(): %s: %s", "Mensagem não enviada", err)
		}

		json.Unmarshal(reqBody, &novaMensagem)

		fmt.Println("Titulo: ", novaMensagem.Titulo)

		if novaMensagem.Titulo != "" && novaMensagem.Texto != "" {
			var mensagemModel = db.Collection("mensagem")
			var interf models.Metodos

			interf = novaMensagem

			//err := interf.Criar(mensagemModel)
			if err := interf.Criar(mensagemModel); err != nil {
				mensagemErro := fmt.Sprintf("EnviarMensagem(): %s: %s", "Erro ao criar a mensagem", err)
				respondError(w, http.StatusInternalServerError, mensagemErro)
				return
			}
		} else {
			mensagem := fmt.Sprint("EnviarMensagem(): Titulo ou Descrição obrigatórios!")
			log.Printf(mensagem)

			respondError(w, http.StatusLengthRequired, mensagem)
			return
		}

		respondJSON(w, http.StatusCreated, novaMensagem)
	}
}
