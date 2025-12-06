package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, router chi.Router) *Server {
	s := &http.Server{

		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{httpServer: s}
}

func (s *Server) Start() {

	// channel to receive OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Server Starting

	go func() {
		log.Printf("Auth service is running on %s\n", s.httpServer.Addr)

		err := s.httpServer.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen: %v\n", err)
		}
	}()

	// wait for interrupt signal
	<-stop
	log.Println("Shutdown signal received")

	s.Shutdown()
}

func (s *Server) Shutdown() {
	// wait for existing connections to finish

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	err := s.httpServer.Shutdown(ctx)

	if err != nil {
		log.Fatalf("Graceful shutdown failed %v\n", err)
	}

	log.Println("Server closed gracefully")
}
