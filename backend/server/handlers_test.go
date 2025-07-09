package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestHandleGetModules(t *testing.T) {
	// Create a test storage
	db := NewDB()

	// Add a module
	db.AddModule(Module{
		Resource: Resource{
			ID:   uuid.New().String(),
			Name: "test-module",
		},
	})

	// Create a server
	server := NewServer(db)

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/modules", nil)
	w := httptest.NewRecorder()

	// Call the handler
	server.getModules(w, req)
	require.EqualValues(t, http.StatusOK, w.Code, "Expected status code 200 OK")

	// Decode the response
	var modules []Module
	err := json.NewDecoder(w.Body).Decode(&modules)
	require.NoError(t, err, "Failed to decode response")

	// Verify the modules
	require.Len(t, modules, 1, "Expected 1 module in response")
	require.Equal(t, "test-module", modules[0].Name, "Expected module name to be 'test-module'")
}

func TestHandleGetTemplates(t *testing.T) {
	db := NewDB()

	// Add a template
	db.AddTemplate(Template{
		Resource: Resource{
			ID:   uuid.New().String(),
			Name: "test-template",
		},
	})

	// Create a server
	server := NewServer(db)

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/templates", nil)
	w := httptest.NewRecorder()

	// Call the handler
	server.getTemplates(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Decode the response
	var templates []Template
	if err := json.NewDecoder(w.Body).Decode(&templates); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Verify the templates
	if len(templates) != 1 {
		t.Errorf("Expected 1 template, got %d", len(templates))
	}

	if templates[0].Name != "test-template" {
		t.Errorf("Expected template name to be 'test-template', got '%s'", templates[0].Name)
	}
}

func TestHandleDeleteModule(t *testing.T) {
	// Create a test storage
	storage := NewDB()

	// Add a module
	module := Module{
		Resource: Resource{
			ID:   uuid.New().String(),
			Name: "test-module",
		},
	}
	storage.AddModule(module)

	// Create a server
	server := NewServer(storage)

	// Create a test HTTP router to use URL params
	r := httptest.NewRequest(http.MethodDelete, "/modules/"+module.ID, nil)
	w := httptest.NewRecorder()

	// Create a router context with URL params
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", module.ID)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	// Call the handler
	server.deleteModule(w, r)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Verify the module was deleted
	modules := storage.GetModules("")
	if len(modules) != 0 {
		t.Errorf("Expected 0 modules, got %d", len(modules))
	}
}
