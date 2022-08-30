package api

import (
	"encoding/json"
	"errors"
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

func (s *server) handleUserDelete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["id"]
		idInt, err := strconv.Atoi(id)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if id == "" || idInt == 0 {
			s.error(w, r, http.StatusBadRequest, errors.New("bad id"))
			return
		}

		err = s.store.User().Delete(idInt)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w,r,http.StatusOK, "Ok")

	}
}

func (s *server) handleUserUpdate() http.HandlerFunc {

	type fields struct {
		UserName string `json:"username,omitempty"`
		Email string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
		RoleId int `json:"role_id,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["id"]
		idInt, err := strconv.Atoi(id)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if id == "" || idInt == 0 {
			s.error(w, r, http.StatusBadRequest, errors.New("bad id"))
			return
		}

		fld := &fields{}
		if err := json.NewDecoder(r.Body).Decode(fld); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().GetById(idInt)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if u.Username == "" {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if fld.RoleId != 0 {u.Role = models.Role{ID: fld.RoleId}}
		if fld.Email != "" {u.Email = fld.Email}
		if fld.Password != "" {u.Password = fld.Password}
		if fld.UserName != "" {u.Username = fld.UserName}

		if err := s.store.User().Update(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.HashPassword = ""
		s.respond(w,r,http.StatusOK, u)

	}
}
