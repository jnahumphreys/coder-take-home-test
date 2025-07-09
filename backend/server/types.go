package server

// OperatingSystem represents the supported operating systems
type OperatingSystem string

// Source represents the source of a resource
type Source string

// Constants for OperatingSystem
const (
	Windows OperatingSystem = "Windows"
	Linux   OperatingSystem = "Linux"
	MacOS   OperatingSystem = "MacOS"
)

// Constants for Source
const (
	Partner  Source = "Partner"
	Official Source = "Official"
)

// Resource is the base struct for both Module and Template
type Resource struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Logo            string          `json:"logo"`
	Contributor     string          `json:"contributor"`
	OperatingSystem OperatingSystem `json:"operating_system"`
	Source          Source          `json:"source"`
	CustomTags      []string        `json:"custom_tags"`
}

// Module represents a Coder module resource
type Module struct {
	Resource
}

// Template represents a Coder template resource
type Template struct {
	Resource
}

// UpdateEvent represents an event for the SSE endpoint
type UpdateEvent struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
