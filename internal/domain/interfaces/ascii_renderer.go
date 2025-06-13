package interfaces

import "ccusage-rainbow/internal/domain/entities"

// FontSize represents the size of ASCII art characters
type FontSize int

const (
	FontSizeSmall  FontSize = 0 // 5x7
	FontSizeMedium FontSize = 1 // 7x9
	FontSizeLarge  FontSize = 2 // 10x13
)

// ASCIIRenderer defines the interface for rendering text as ASCII art
type ASCIIRenderer interface {
	// RenderPlain renders text as plain ASCII art without colors
	RenderPlain(text *entities.Text) (string, error)

	// RenderPlainWithSize renders text with specified font size
	RenderPlainWithSize(text *entities.Text, size FontSize) (string, error)

	// GetDisplayWidth calculates the actual display width of rendered text
	GetDisplayWidth(rendered string) int

	// GetDisplayWidthWithSize calculates display width for specific font size
	GetDisplayWidthWithSize(text *entities.Text, size FontSize) (int, error)
}
