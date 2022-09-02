package api

import (
	"encoding/json"
	"errors"
	"goApiAuth/go/internal/models"
	"net/http"
	"github.com/gorilla/mux"
)

func (s *server) handleTokenCreate() http.HandlerFunc {

	type request struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		t := &models.Token{
			User: models.User{
				Username: req.UserName,
				Password: req.Password,
			},
		}

		if err := s.store.Token().Create(t); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		t.User.HashPassword = ""
		t.User.Password = ""

		s.respond(w,r,http.StatusOK, t)

	}
}

func (s *server) handleTokenCheck() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {


		params := mux.Vars(r)
		tokenStr := params["token"]

		if tokenStr == "" {
			s.error(w, r, http.StatusBadRequest, errors.New("bad token"))
			return

		}

		token, err := s.store.Token().Check(tokenStr)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		s.respond(w,r,http.StatusOK, token)

	}
}
