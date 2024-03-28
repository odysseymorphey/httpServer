package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/nats-io/stan.go"
	"github.com/odysseymorphey/httpServer/internal/database"
	"github.com/odysseymorphey/httpServer/internal/model"
	"log"
	"net/http"
)

type Server struct {
	router *chi.Mux
	server *http.Server
	db     *database.DB
	cache  map[string]model.Order
	sc     stan.Conn
	sub    stan.Subscription
}

func NewServer() (*Server, error) {
	mux := chi.NewRouter()

	db, err := database.NewDB()
	if err != nil {
		return nil, err
	}

	server := &Server{
		router: mux,
		db:     db,
		cache:  make(map[string]model.Order),
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("TESTING"))
	})

	return server, nil
}

func (s *Server) Run() {
	s.server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: s.router,
	}

	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
