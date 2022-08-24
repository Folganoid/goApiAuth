package api

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"goApiAuth/go/internal/store"
	"net/http"
)

type server struct {
	router *mux.Router
	store store.Store
}

func newServer(store store.Store) *server {
	s := &server {
		router: mux.NewRouter(),
		store: store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/test", s.handleTest()).Methods("GET")
}

func (s *server) handleTest() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Ok")
	}
}