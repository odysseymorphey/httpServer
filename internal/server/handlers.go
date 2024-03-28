package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("ui/index.html"))

func (s *Server) mock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.cache)
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("order_id")

	tmpl.Execute(w, id)

	w.Write([]byte(id))
}

func (s *Server) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "order_id")
	item, ok := s.cache[id]
	if !ok {
		w.Write([]byte("Something went wrong"))
		return
	}

	b, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		return
	}

	w.Write(b)

}
