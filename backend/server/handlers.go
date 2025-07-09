package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// getModules returns a list of modules, optionally filtered by name
func (s *Server) getModules(w http.ResponseWriter, r *http.Request) {
	nameFilter := r.URL.Query().Get("name")
	modules := s.db.GetModules(nameFilter)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(modules); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// getTemplates returns a list of templates, optionally filtered by name
func (s *Server) getTemplates(w http.ResponseWriter, r *http.Request) {
	nameFilter := r.URL.Query().Get("name")
	templates := s.db.GetTemplates(nameFilter)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(templates); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// autocompleteModules returns module names that match a prefix
func (s *Server) autocompleteModules(w http.ResponseWriter, r *http.Request) {
	prefix := r.URL.Query().Get("prefix")
	suggestions := s.db.GetModuleSuggestions(prefix)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(suggestions); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// autocompleteTemplates returns template names that match a prefix
func (s *Server) autocompleteTemplates(w http.ResponseWriter, r *http.Request) {
	prefix := r.URL.Query().Get("prefix")
	suggestions := s.db.GetTemplateSuggestions(prefix)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(suggestions); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// deleteModule deletes a module by ID
func (s *Server) deleteModule(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Module ID is required", http.StatusBadRequest)
		return
	}

	if deleted := s.db.DeleteModule(id); !deleted {
		http.Error(w, "Module not found", http.StatusNotFound)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Module deleted"})
}

// deleteTemplate deletes a template by ID
func (s *Server) deleteTemplate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Template ID is required", http.StatusBadRequest)
		return
	}

	if deleted := s.db.DeleteTemplate(id); !deleted {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Template deleted"})
}

// sendEvents creates a Server-Sent Events stream for live updates
func (s *Server) streamEvents(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the flusher to ensure headers and events are sent immediately
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}
	flusher.Flush()

	// Get the SSE channel from storage
	updatesCh := s.db.Updates()

	// Listen for client disconnect
	clientDisconnect := r.Context().Done()

	// Send initial event to establish connection
	fmt.Fprintf(w, "event: connected\ndata: {\"status\":\"connected\"}\n\n")
	flusher.Flush()

	// Event loop
	for {
		select {
		case update, ok := <-updatesCh:
			// Channel was closed
			if !ok {
				return
			}

			// Marshal the update data
			data, err := json.Marshal(update)
			if err != nil {
				continue
			}

			// Write the event to the response
			fmt.Fprintf(w, "event: message\ndata: %s\n\n", data)
			flusher.Flush()

		case <-clientDisconnect:
			// Client disconnected
			return
		}
	}
}
