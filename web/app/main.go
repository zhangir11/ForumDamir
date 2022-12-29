package main

import (
	"fmt"
	"forum/internal/router"
	"log"
	"net/http"
	"time"
)

func main() {
	// create handlers
	handlers, err := router.NewMainHandler()
	if err != nil {
		log.Fatal(err)
	}
	server := new(Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("ERROR: %s\n", err)
	}
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second, // return cont.Done() ifup time limit
		WriteTimeout:   5 * time.Second, // return cont.Done() ifup time limit
	}

	log.Printf("Server runs on http://localhost%s\n", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	return fmt.Errorf("Server.Run: %w", err)
}
