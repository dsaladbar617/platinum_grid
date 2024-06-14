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

	// SHEET ROUTES
	mux.HandleFunc("POST /sheet", s.addSheet)
	mux.HandleFunc("PUT /edit_sheet/{id}", s.updateSheet)
	mux.HandleFunc("GET /get_sheet/{id}", s.getSheetByID)

	// USER ROUTES
	mux.HandleFunc("POST /add_user", s.addUser)

	mux.HandleFunc("/health", s.healthHandler)

	root := http.NewServeMux()
	root.Handle("/api/", http.StripPrefix("/api", mux))

	return standard.Then(root)
}
