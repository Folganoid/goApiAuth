package api

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"goApiAuth/go/internal/store"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type server struct {
	router *mux.Router
	store store.Store
	logger *log.Logger
}

func newServer(store store.Store, logger *log.Logger) *server {
	s := &server {
		router: mux.NewRouter(),
		store: store,
		logger: logger,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/test", s.handleTest()).Methods("GET")

	//user
	s.router.HandleFunc("/user", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/user/id/{id}", s.handleUserGetById()).Methods("GET")
	s.router.HandleFunc("/user/id/{id}", s.handleUserUpdate()).Methods("PUT")
	s.router.HandleFunc("/user/id/{id}", s.handleUserDelete()).Methods("DELETE")

	//token
	s.router.HandleFunc("/token", s.handleTokenCreate()).Methods("POST")
	s.router.HandleFunc("/token/{token}", s.handleTokenCheck()).Methods("GET")

	//role
	s.router.HandleFunc("/role/{id}", s.handleRoleGet()).Methods("GET")

	s.logger.Info("API started")
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(log.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level log.Level
		switch {
		case rw.code >= 500:
			level = log.ErrorLevel
		case rw.code >= 400:
			level = log.WarnLevel
		default:
			level = log.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().String(),
		)
	})
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	s.logger.Debugf("response body: %v", data)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}