package server

import "net/http"

func (s *Server) mock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TESTING"))
}
