package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marceloagmelo/go-rabbitmq-send/app/handler"
	"github.com/marceloagmelo/go-rabbitmq-send/config"
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

const (
	staticDir = "/static/"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     db.Database
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	configuracao, err := mysql.ParseURL(dbURI)

	a.DB, err = mysql.Open(configuracao)
	if err != nil {
		log.Fatal(err.Error())
	}

	a.Router = mux.NewRouter()
	a.Router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	a.Get("/", a.handleDBRequest(handler.Home))
	a.Get("/health", a.handleDBRequest(handler.Health))
	a.Get("/v1/new", a.handleRequest(handler.New))
	a.Post("/v1/insert", a.handleDBRequest(handler.Insert))
	a.Get("/v1/mensagens", a.handleDBRequest(handler.TodasMensagens))
	a.Get("/v1/mensagem/{id}", a.handleDBRequest(handler.UmaMensagem))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

//RequestHandlerFunction função handler
type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

//RequestHandlerDBFunction função handler
type RequestHandlerDBFunction func(db db.Database, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

func (a *App) handleDBRequest(handler RequestHandlerDBFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
