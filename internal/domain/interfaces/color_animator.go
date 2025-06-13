package interfaces

import "ccusage-rainbow/internal/domain/entities"

// ColorAnimator defines the interface for applying animated colors to text
type ColorAnimator interface {
	// ApplyRainbowColors applies animated rainbow colors to ASCII art
	ApplyRainbowColors(asciiArt string, animation *entities.RainbowAnimation) string
}
