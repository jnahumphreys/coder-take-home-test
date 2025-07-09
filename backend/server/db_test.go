package server

import (
	"testing"

	"github.com/google/uuid"
)

func TestStorage_AddGetModules(t *testing.T) {
	db := NewDB()

	// Add a module
	module := Module{
		Resource: Resource{
			ID:              uuid.New().String(),
			Name:            "test-module",
			Description:     "Test module description",
			Logo:            "https://example.com/logo.png",
			Contributor:     "Test Contributor",
			OperatingSystem: Linux,
			Source:          Official,
			CustomTags:      []string{"test", "module"},
		},
	}

	db.AddModule(module)

	// Get all modules
	modules := db.GetModules("")
	if len(modules) != 1 {
		t.Errorf("Expected 1 module, got %d", len(modules))
	}

	// Check the module data
	if modules[0].Name != "test-module" {
		t.Errorf("Expected module name to be 'test-module', got '%s'", modules[0].Name)
	}
}

func TestStorage_DeleteModule(t *testing.T) {
	db := NewDB()

	// Add a module
	module := Module{
		Resource: Resource{
			ID:              uuid.New().String(),
			Name:            "test-module",
			Description:     "Test module description",
			Logo:            "https://example.com/logo.png",
			Contributor:     "Test Contributor",
			OperatingSystem: Linux,
			Source:          Official,
			CustomTags:      []string{"test", "module"},
		},
	}

	db.AddModule(module)

	// Delete the module by ID
	deleted := db.DeleteModule(module.ID)
	if !deleted {
		t.Error("Expected module to be deleted")
	}

	// Check that the module is gone
	modules := db.GetModules("")
	if len(modules) != 0 {
		t.Errorf("Expected 0 modules, got %d", len(modules))
	}
}

func TestStorage_FilterModulesByName(t *testing.T) {
	db := NewDB()

	// Add modules
	db.AddModule(Module{
		Resource: Resource{
			ID:   uuid.New().String(),
			Name: "first-module",
		},
	})
	db.AddModule(Module{
		Resource: Resource{
			ID:   uuid.New().String(),
			Name: "second-module",
		},
	})

	// Filter by name
	modules := db.GetModules("first")
	if len(modules) != 1 {
		t.Errorf("Expected 1 module, got %d", len(modules))
	}

	if len(modules) > 0 && modules[0].Name != "first-module" {
		t.Errorf("Expected module name to be 'first-module', got '%s'", modules[0].Name)
	}
}

func TestStorage_ModuleSuggestions(t *testing.T) {
	db := NewDB()

	// Add modules
	db.AddModule(Module{
		Resource: Resource{
			ID:   uuid.New().String(),
			Name: "app-module",
		},
	})
	db.AddModule(Module{
		Resource: Resource{
			ID:   uuid.New().String(),
			Name: "awesome-module",
		},
	})
	db.AddModule(Module{
		Resource: Resource{
			ID:   uuid.New().String(),
			Name: "basic-module",
		},
	})

	// Get suggestions
	suggestions := db.GetModuleSuggestions("a")
	if len(suggestions) != 2 {
		t.Errorf("Expected 2 suggestions, got %d", len(suggestions))
	}
}
