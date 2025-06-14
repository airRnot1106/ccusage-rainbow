package ascii

import (
	"ccusage-rainbow/internal/domain/entities"
	"ccusage-rainbow/internal/domain/interfaces"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Renderer implements the ASCIIRenderer interface
type Renderer struct {
	smallPatterns  map[rune][]string
	mediumPatterns map[rune][]string
	largePatterns  map[rune][]string
}

// NewRenderer creates a new ASCII renderer
func NewRenderer() *Renderer {
	return &Renderer{
		smallPatterns:  getSmallLetterPatterns(),
		mediumPatterns: getMediumLetterPatterns(),
		largePatterns:  getLargeLetterPatterns(),
	}
}

// RenderPlain renders text as plain ASCII art without colors
func (r *Renderer) RenderPlain(text *entities.Text) (string, error) {
	content := strings.ToUpper(text.Content)

	var result []string
	for i := 0; i < 7; i++ { // 7 rows per character
		var line string
		for j, char := range content {
			if patterns, ok := r.mediumPatterns[char]; ok {
				line += patterns[i]
			} else {
				line += "         " // Default spacing for unknown characters (9 chars)
			}
			// Only add space between characters, not after the last one
			if j < len(content)-1 {
				line += " "
			}
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n"), nil
}

// RenderPlainWithSize renders text with specified font size
func (r *Renderer) RenderPlainWithSize(text *entities.Text, size interfaces.FontSize) (string, error) {
	content := strings.ToUpper(text.Content)

	var patterns map[rune][]string
	var rows int
	var defaultSpacing string

	switch size {
	case interfaces.FontSizeSmall:
		patterns = r.smallPatterns
		rows = 5
		defaultSpacing = "       " // 7 chars
	case interfaces.FontSizeMedium:
		patterns = r.mediumPatterns
		rows = 7
		defaultSpacing = "         " // 9 chars
	case interfaces.FontSizeLarge:
		patterns = r.largePatterns
		rows = 10
		defaultSpacing = "             " // 13 chars
	default:
		patterns = r.mediumPatterns
		rows = 7
		defaultSpacing = "         "
	}

	var result []string
	for i := 0; i < rows; i++ {
		var line string
		for j, char := range content {
			if charPatterns, ok := patterns[char]; ok {
				line += charPatterns[i]
			} else {
				line += defaultSpacing
			}
			// Only add space between characters, not after the last one
			if j < len(content)-1 {
				nextChar := rune(content[j+1])
				currentChar := char

				// Use smaller spacing around decimal points
				if currentChar == '.' || nextChar == '.' {
					switch size {
					case interfaces.FontSizeSmall:
						line += " " // 1 space around decimal point for small font
					case interfaces.FontSizeMedium:
						line += "  " // 2 spaces around decimal point for medium font
					case interfaces.FontSizeLarge:
						line += "  " // 2 spaces around decimal point for large font
					default:
						line += "  " // default 2 spaces around decimal point
					}
				} else {
					switch size {
					case interfaces.FontSizeSmall:
						line += "  " // 2 spaces for small font
					case interfaces.FontSizeMedium:
						line += "   " // 3 spaces for medium font
					case interfaces.FontSizeLarge:
						line += "    " // 4 spaces for large font
					default:
						line += "   " // default 3 spaces
					}
				}
			}
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n"), nil
}

// GetDisplayWidth calculates the actual display width of rendered text
func (r *Renderer) GetDisplayWidth(rendered string) int {
	lines := strings.Split(rendered, "\n")
	maxWidth := 0

	for _, line := range lines {
		if line != "" {
			width := lipgloss.Width(line)
			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	return maxWidth
}

// GetDisplayWidthWithSize calculates display width for specific font size
func (r *Renderer) GetDisplayWidthWithSize(text *entities.Text, size interfaces.FontSize) (int, error) {
	rendered, err := r.RenderPlainWithSize(text, size)
	if err != nil {
		return 0, err
	}
	return r.GetDisplayWidth(rendered), nil
}

// getMediumLetterPatterns returns the medium-size ASCII art patterns (7x9)
func getMediumLetterPatterns() map[rune][]string {
	return map[rune][]string{
		'0': {
			" ███████ ",
			"███   ███",
			"███   ███",
			"███   ███",
			"███   ███",
			"███   ███",
			" ███████ ",
		},
		'1': {
			"   ███   ",
			" █████   ",
			"   ███   ",
			"   ███   ",
			"   ███   ",
			"   ███   ",
			" ███████ ",
		},
		'2': {
			" ███████ ",
			"███   ███",
			"      ███",
			" ███████ ",
			"███      ",
			"███      ",
			"█████████",
		},
		'3': {
			" ███████ ",
			"███   ███",
			"      ███",
			"   █████ ",
			"      ███",
			"███   ███",
			" ███████ ",
		},
		'4': {
			"███   ███",
			"███   ███",
			"███   ███",
			"█████████",
			"      ███",
			"      ███",
			"      ███",
		},
		'5': {
			"█████████",
			"███      ",
			"███      ",
			"████████ ",
			"      ███",
			"███   ███",
			" ███████ ",
		},
		'6': {
			" ███████ ",
			"███   ███",
			"███      ",
			"████████ ",
			"███   ███",
			"███   ███",
			" ███████ ",
		},
		'7': {
			"█████████",
			"      ███",
			"     ███ ",
			"    ███  ",
			"   ███   ",
			"  ███    ",
			" ███     ",
		},
		'8': {
			" ███████ ",
			"███   ███",
			"███   ███",
			" ███████ ",
			"███   ███",
			"███   ███",
			" ███████ ",
		},
		'9': {
			" ███████ ",
			"███   ███",
			"███   ███",
			" ████████",
			"      ███",
			"███   ███",
			" ███████ ",
		},
		'.': {
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			" ███  ",
			" ███  ",
		},
		'$': {
			"    ███  ",
			" ███████ ",
			"███ ███  ",
			" ███████ ",
			"  ███ ███",
			" ███████ ",
			"   ███   ",
		},
		' ': {
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
		},
		'E': {
			"█████████",
			"███      ",
			"███      ",
			"███████  ",
			"███      ",
			"███      ",
			"█████████",
		},
		'R': {
			"████████ ",
			"███   ███",
			"███   ███",
			"████████ ",
			"███  ███ ",
			"███   ███",
			"███   ███",
		},
		'O': {
			" ███████ ",
			"███   ███",
			"███   ███",
			"███   ███",
			"███   ███",
			"███   ███",
			" ███████ ",
		},
		'L': {
			"███      ",
			"███      ",
			"███      ",
			"███      ",
			"███      ",
			"███      ",
			"█████████",
		},
		'H': {
			"███   ███",
			"███   ███",
			"███   ███",
			"█████████",
			"███   ███",
			"███   ███",
			"███   ███",
		},
	}
}

// getSmallLetterPatterns returns small-size ASCII art patterns (5x7)
func getSmallLetterPatterns() map[rune][]string {
	return map[rune][]string{
		'$': {"  ███  ", " █████ ", "███    ", " █████ ", "   ███ "},
		'0': {" █████ ", "██   ██", "██   ██", "██   ██", " █████ "},
		'1': {"  ███  ", " ████  ", "  ███  ", "  ███  ", "███████"},
		'2': {" █████ ", "     ██", " █████ ", "██     ", "███████"},
		'3': {" █████ ", "     ██", "  ████ ", "     ██", " █████ "},
		'4': {"██  ██ ", "██  ██ ", "███████", "    ██ ", "    ██ "},
		'5': {"███████", "██     ", "██████ ", "     ██", "██████ "},
		'6': {" █████ ", "██     ", "██████ ", "██   ██", " █████ "},
		'7': {"███████", "     ██", "    ██ ", "   ██  ", "  ██   "},
		'8': {" █████ ", "██   ██", " █████ ", "██   ██", " █████ "},
		'9': {" █████ ", "██   ██", " ██████", "     ██", " █████ "},
		'.': {"       ", "       ", "       ", "  ██   ", "  ██   "},
		' ': {"       ", "       ", "       ", "       ", "       "},
		'E': {"███████", "██     ", "██████ ", "██     ", "███████"},
		'R': {"██████ ", "██   ██", "██████ ", "██  ██ ", "██   ██"},
		'O': {" █████ ", "██   ██", "██   ██", "██   ██", " █████ "},
		'L': {"██     ", "██     ", "██     ", "██     ", "███████"},
		'H': {"██   ██", "██   ██", "███████", "██   ██", "██   ██"},
	}
}

// getLargeLetterPatterns returns large-size ASCII art patterns (10x13)
func getLargeLetterPatterns() map[rune][]string {
	return map[rune][]string{
		'$': {
			"     ████     ",
			"  ███████████ ",
			" ████ ███     ",
			"████  ████    ",
			" ███████████  ",
			"  ███████████ ",
			"     ████ ████",
			"████████  ████",
			" ███████████  ",
			"     ████     ",
		},
		'0': {
			"  ██████████  ",
			" ████    ████ ",
			"████      ████",
			"████      ████",
			"████      ████",
			"████      ████",
			"████      ████",
			"████      ████",
			" ████    ████ ",
			"  ██████████  ",
		},
		'1': {
			"     ████     ",
			"  ███████     ",
			"     ████     ",
			"     ████     ",
			"     ████     ",
			"     ████     ",
			"     ████     ",
			"     ████     ",
			"     ████     ",
			"██████████████",
		},
		'2': {
			"  ███████████ ",
			" ████     ████",
			"          ████",
			"         ████ ",
			"       ████   ",
			"     ████     ",
			"   ████       ",
			" ████         ",
			"████          ",
			"██████████████",
		},
		'3': {
			"  ███████████ ",
			" ████     ████",
			"          ████",
			"          ████",
			"     █████████",
			"          ████",
			"          ████",
			"          ████",
			" ████     ████",
			"  ███████████ ",
		},
		'4': {
			"████      ████",
			"████      ████",
			"████      ████",
			"████      ████",
			"██████████████",
			"          ████",
			"          ████",
			"          ████",
			"          ████",
			"          ████",
		},
		'5': {
			"██████████████",
			"████          ",
			"████          ",
			"████          ",
			"█████████████ ",
			"          ████",
			"          ████",
			"          ████",
			" ████     ████",
			"  ███████████ ",
		},
		'6': {
			"  ███████████ ",
			" ████     ████",
			"████          ",
			"████          ",
			"█████████████ ",
			"████      ████",
			"████      ████",
			"████      ████",
			" ████     ████",
			"  ███████████ ",
		},
		'7': {
			"██████████████",
			"██████████████",
			"          ████",
			"         ████ ",
			"        ████  ",
			"       ████   ",
			"      ████    ",
			"     ████     ",
			"    ████      ",
			"   ████       ",
		},
		'8': {
			"  ██████████  ",
			" ████    ████ ",
			"████      ████",
			" ████    ████ ",
			"  ██████████  ",
			" ████    ████ ",
			"████      ████",
			"████      ████",
			" ████    ████ ",
			"  ██████████  ",
		},
		'9': {
			"  ██████████  ",
			" ████    ████ ",
			"████      ████",
			"████      ████",
			" █████████████",
			"          ████",
			"          ████",
			"          ████",
			" ████     ███ ",
			"  ██████████  ",
		},
		'.': {
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			"         ",
			" ██████  ",
			" ██████  ",
			" ██████  ",
		},
		' ': {
			"              ",
			"              ",
			"              ",
			"              ",
			"              ",
			"              ",
			"              ",
			"              ",
			"              ",
			"              ",
		},
		'E': {
			"█████████████",
			"████         ",
			"████         ",
			"████         ",
			"█████████    ",
			"████         ",
			"████         ",
			"████         ",
			"████         ",
			"█████████████",
		},
		'R': {
			"████████████ ",
			"████     ████",
			"████     ████",
			"████     ████",
			"████████████ ",
			"████   ████  ",
			"████    ████ ",
			"████     ████",
			"████     ████",
			"████     ████",
		},
		'O': {
			"  █████████  ",
			" ████   ████ ",
			"████     ████",
			"████     ████",
			"████     ████",
			"████     ████",
			"████     ████",
			"████     ████",
			" ████   ████ ",
			"  █████████  ",
		},
		'L': {
			"████         ",
			"████         ",
			"████         ",
			"████         ",
			"████         ",
			"████         ",
			"████         ",
			"████         ",
			"████         ",
			"█████████████",
		},
		'H': {
			"████     ████",
			"████     ████",
			"████     ████",
			"████     ████",
			"█████████████",
			"████     ████",
			"████     ████",
			"████     ████",
			"████     ████",
			"████     ████",
		},
	}
}
