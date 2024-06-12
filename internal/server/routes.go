package server

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/rs/cors"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	// c := cors.Default()
	standard := alice.New(cors.Default().Handler)

	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("POST /sheet", s.addSheet)

	mux.HandleFunc("/health", s.healthHandler)

	root := http.NewServeMux()
	root.Handle("/api/", http.StripPrefix("/api", mux))

	return standard.Then(root)
}
