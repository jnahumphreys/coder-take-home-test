package server

import (
	"strings"
	"sync"
)

// DB handles storing and retrieving modules and templates in memory
type DB struct {
	modules   []Module
	templates []Template
	mu        sync.RWMutex
	updates   chan UpdateEvent
	closed    bool
}

// NewDB creates a new memory db instance
func NewDB() *DB {
	return &DB{
		modules:   []Module{},
		templates: []Template{},
		updates:   make(chan UpdateEvent, 100), // Buffered channel to prevent blocking
	}
}

// AddModule adds a new module to storage and broadcasts an update event
func (s *DB) AddModule(module Module) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.modules = append(s.modules, module)

	// Send update event
	select {
	case s.updates <- UpdateEvent{Type: "module_added", Data: module}:
	default:
		// Channel is full, don't block
	}
}

// AddTemplate adds a new template to storage and broadcasts an update event
func (s *DB) AddTemplate(template Template) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.templates = append(s.templates, template)

	// Send update event
	select {
	case s.updates <- UpdateEvent{Type: "template_added", Data: template}:
	default:
		// Channel is full, don't block
	}
}

// GetModules returns all modules, optionally filtered by name
func (s *DB) GetModules(nameFilter string) []Module {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if nameFilter == "" {
		// Return a copy to prevent modification of internal state
		result := make([]Module, len(s.modules))
		copy(result, s.modules)
		return result
	}

	// Filter by name
	var filtered []Module
	for _, m := range s.modules {
		if strings.Contains(strings.ToLower(m.Name), strings.ToLower(nameFilter)) {
			filtered = append(filtered, m)
		}
	}
	return filtered
}

// GetTemplates returns all templates, optionally filtered by name
func (s *DB) GetTemplates(nameFilter string) []Template {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if nameFilter == "" {
		// Return a copy to prevent modification of internal state
		result := make([]Template, len(s.templates))
		copy(result, s.templates)
		return result
	}

	// Filter by name
	var filtered []Template
	for _, t := range s.templates {
		if strings.Contains(strings.ToLower(t.Name), strings.ToLower(nameFilter)) {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

// DeleteModule removes a module by ID
func (s *DB) DeleteModule(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, m := range s.modules {
		if strings.EqualFold(m.ID, id) {
			// Remove the module by replacing it with the last one
			// and shrinking the slice (faster than creating a new slice)
			s.modules[i] = s.modules[len(s.modules)-1]
			s.modules = s.modules[:len(s.modules)-1]

			// Send update event
			select {
			case s.updates <- UpdateEvent{Type: "module_deleted", Data: m}:
			default:
				// Channel is full, don't block
			}

			return true
		}
	}
	return false
}

// DeleteTemplate removes a template by ID
func (s *DB) DeleteTemplate(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, t := range s.templates {
		if strings.EqualFold(t.ID, id) {
			// Remove the template by replacing it with the last one
			// and shrinking the slice (faster than creating a new slice)
			s.templates[i] = s.templates[len(s.templates)-1]
			s.templates = s.templates[:len(s.templates)-1]

			// Send update event
			select {
			case s.updates <- UpdateEvent{Type: "template_deleted", Data: t}:
			default:
				// Channel is full, don't block
			}

			return true
		}
	}
	return false
}

// GetModuleSuggestions returns module names that start with the given prefix
func (s *DB) GetModuleSuggestions(prefix string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	prefix = strings.ToLower(prefix)
	var suggestions []string

	for _, m := range s.modules {
		if strings.HasPrefix(strings.ToLower(m.Name), prefix) {
			suggestions = append(suggestions, m.Name)
		}
	}
	return suggestions
}

// GetTemplateSuggestions returns template names that start with the given prefix
func (s *DB) GetTemplateSuggestions(prefix string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	prefix = strings.ToLower(prefix)
	var suggestions []string

	for _, t := range s.templates {
		if strings.HasPrefix(strings.ToLower(t.Name), prefix) {
			suggestions = append(suggestions, t.Name)
		}
	}
	return suggestions
}

// Updates returns a read-only channel for receiving update events
func (s *DB) Updates() <-chan UpdateEvent {
	return s.updates
}

// Close closes the updates channel
func (s *DB) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return nil // Already closed, no-op
	}

	s.closed = true
	close(s.updates)
	return nil
}
