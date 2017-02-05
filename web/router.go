package web

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"log"

	"github.com/gorilla/mux"
)

type WebServer struct {
	router *mux.Router
}

func New() *WebServer {
	return &WebServer{
		router: mux.NewRouter(),
	}
}

func (webServer *WebServer) InitHandlers() {

	// index
	// webServer.router.HandleFunc("/", webServer.IndexHandler)

	// static content
	webServer.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/static/"))))
}

func (webServer *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := handlers.CombinedLoggingHandler(os.Stdout, webServer.router)
	router.ServeHTTP(w, r)
}

func (webServer *WebServer) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, webServer))
}
