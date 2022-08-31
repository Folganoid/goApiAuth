package api

import (
	"encoding/json"
	"goApiAuth/go/internal/models"
	"net/http"
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
