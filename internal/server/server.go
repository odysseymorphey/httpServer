package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/stan.go"
	"github.com/odysseymorphey/httpServer/internal/database"
	"github.com/odysseymorphey/httpServer/internal/model"
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

	db := database.NewDB()

	server := &Server{
		router: mux,
		db:     db,
		cache:  make(map[string]model.Order),
	}

	return server, nil
}

func (s *Server) Run() error {
	err := s.setCache()
	if err != nil {
		return err
	}

	err = s.connectToStream()
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		select {
		case <-sig:
		}

		s.Stop()
		os.Exit(0)
	}()

	if err = s.beginRouting(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.db.Close()
	s.sub.Unsubscribe()
	s.sc.Close()
}

func (s *Server) beginRouting() error {
	s.server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: s.router,
	}

	s.router.HandleFunc("/", s.mock)
	s.router.HandleFunc("/tmpl", s.testHandler)
	s.router.Get("/order/{order_id}", s.handleGetOrder)

	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setCache() error {
	orders := make([]model.Order, 0)

	err := s.db.DB.Model(&orders).Select()
	if err != nil {
		return nil
	}

	for _, order := range orders {
		s.cache[order.OrderUid] = order
	}

	return nil
}

func (s *Server) connectToStream() error {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("localhost:4222"))
	if err != nil {
		return err
	}

	sub, err := sc.Subscribe("addNewOrder", s.handleRequest)
	if err != nil {
		return err
	}

	s.sc, s.sub = sc, sub

	return nil
}

func (s *Server) handleRequest(m *stan.Msg) {
	data := model.Order{}
	err := json.Unmarshal(m.Data, &data)
	if err != nil {
		return
	}

	if ok := s.addToCache(data); ok {
		log.Print("Added to cache")
		s.db.AddOrder(data)
	}
}

func (s *Server) addToCache(data model.Order) bool {
	_, ok := s.cache[data.OrderUid]
	if ok {
		return false
	}

	s.cache[data.OrderUid] = data
	for key := range s.cache {
		fmt.Printf("%s ", key)
	}
	fmt.Println()
	return true
}
