package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func (s *server) handleTest() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Ok")
	}
}

func (s *server) handleRoleGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["id"]
		idInt, err := strconv.Atoi(id)

		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		role, err := s.store.Role().GetById(idInt)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w,r,http.StatusOK, role)

	}
}
