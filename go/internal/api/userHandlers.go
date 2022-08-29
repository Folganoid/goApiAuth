package api

import (
	"encoding/json"
	"goApiAuth/go/internal/models"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func (s *server) handleUsersCreate() http.HandlerFunc {

	type request struct {
		UserName string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Username: req.UserName,
			Email: req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.HashPassword = ""

		s.respond(w,r,http.StatusOK, u)

	}
}

func (s *server) handleUserGetById() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["id"]
		idInt, err := strconv.Atoi(id)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		role, err := s.store.User().GetById(idInt)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w,r,http.StatusOK, role)

	}
}

func (s *server) handleUserGetByEmail() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		email := params["email"]

		role, err := s.store.User().GetByEmail(email)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w,r,http.StatusOK, role)

	}
}
