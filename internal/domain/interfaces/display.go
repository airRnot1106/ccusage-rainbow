package interfaces

// DisplayDimensions represents terminal display dimensions
type DisplayDimensions struct {
	Width  int
	Height int
}

// Display defines the interface for terminal display operations
type Display interface {
	// GetDimensions returns the current terminal dimensions
	GetDimensions() DisplayDimensions

	// Render displays the content on screen
	Render(content string) error

	// Clear clears the display
	Clear() error
}
