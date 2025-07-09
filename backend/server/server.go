package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server represents the HTTP server
type Server struct {
	router *chi.Mux
	db     *DB
}

// NewServer creates a new server instance
func NewServer(storage *DB) *Server {
	s := &Server{
		db: storage,
	}

	// Setup router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/modules", s.getModules)
	r.Get("/templates", s.getTemplates)
	r.Get("/autocomplete/modules", s.autocompleteModules)
	r.Get("/autocomplete/templates", s.autocompleteTemplates)
	r.Delete("/modules/{id}", s.deleteModule)
	r.Delete("/templates/{id}", s.deleteTemplate)
	r.Get("/events", s.streamEvents)

	// Set the router
	s.router = r

	return s
}

// setupRoutes configures the HTTP routes
func (s *Server) setupRoutes() {
	// Middleware
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	// Routes
	s.router.Get("/modules", s.getModules)
	s.router.Get("/templates", s.getTemplates)
	s.router.Get("/autocomplete/modules", s.autocompleteModules)
	s.router.Get("/autocomplete/templates", s.autocompleteTemplates)
	s.router.Delete("/modules/{id}", s.deleteModule)
	s.router.Delete("/templates/{id}", s.deleteTemplate)
	s.router.Get("/events", s.streamEvents)
}

// ServeHTTP implements the http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Listen starts the HTTP server on the given address
func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
