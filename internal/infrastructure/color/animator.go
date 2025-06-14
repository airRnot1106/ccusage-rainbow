package color

import (
	"ccusage-rainbow/internal/domain/entities"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Animator implements the ColorAnimator interface
type Animator struct {
	rainbowColors []string
}

// NewAnimator creates a new color animator
func NewAnimator() *Animator {
	return &Animator{
		rainbowColors: []string{
			"#FF0000", // Red
			"#FF8000", // Orange
			"#FFFF00", // Yellow
			"#00FF00", // Green
			"#0080FF", // Blue
			"#4000FF", // Indigo
			"#8000FF", // Violet
		},
	}
}

// ApplyRainbowColors applies animated rainbow colors to ASCII art
func (a *Animator) ApplyRainbowColors(asciiArt string, animation *entities.RainbowAnimation) string {
	var result strings.Builder
	colorIndex := animation.GetOffset()

	for _, char := range asciiArt {
		if char == ' ' || char == '\n' {
			result.WriteRune(char)
		} else {
			colorIdx := colorIndex % len(a.rainbowColors)
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(a.rainbowColors[colorIdx]))
			result.WriteString(style.Render(string(char)))
			colorIndex = (colorIndex + 1) % len(a.rainbowColors)
		}
	}

	return result.String()
}
