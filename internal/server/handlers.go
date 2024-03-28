package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *Server) mock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TESTING"))
}

func (s *Server) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, chi.URLParam(r, "order_id"))

}
