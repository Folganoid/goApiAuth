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

	//user
	s.router.HandleFunc("/user", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/user/id/{id}", s.handleUserGetById()).Methods("GET")
	s.router.HandleFunc("/user/id/{id}", s.handleUserUpdate()).Methods("PUT")
	s.router.HandleFunc("/user/id/{id}", s.handleUserDelete()).Methods("DELETE")

	//role
	s.router.HandleFunc("/role/{id}", s.handleRoleGet()).Methods("GET")
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}