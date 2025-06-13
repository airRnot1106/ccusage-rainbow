package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type model struct {
	text        string
	width       int
	height      int
	colorOffset int
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tickMsg:
		m.colorOffset = (m.colorOffset + 1) % 7
		return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}
	return m, nil
}

func (m model) View() string {
	// Handle case when dimensions are not set yet
	if m.width <= 0 || m.height <= 0 {
		return "Loading..."
	}

	// Generate ASCII art without colors first
	plainAsciiArt := generatePlainBigText(m.text)

	// Center the ASCII art on screen
	lines := strings.Split(plainAsciiArt, "\n")
	var centeredLines []string

	// Check if ASCII art fits within screen width (using actual display width)
	maxLineWidth := 0
	for _, line := range lines {
		if line != "" {
			displayWidth := lipgloss.Width(line)
			if displayWidth > maxLineWidth {
				maxLineWidth = displayWidth
			}
		}
	}

	// If text doesn't fit, show fallback message
	if maxLineWidth > m.width {
		fallbackMsg := "Terminal too small"
		padding := (m.width - len(fallbackMsg)) / 2
		if padding < 0 {
			padding = 0
		}
		verticalPadding := m.height / 2
		result := strings.Repeat("\n", verticalPadding) + strings.Repeat(" ", padding) + fallbackMsg

		quitPadding := (m.width - 20) / 2
		if quitPadding < 0 {
			quitPadding = 0
		}
		result += "\n\n" + strings.Repeat(" ", quitPadding) + "Press 'q' to quit"
		return result
	}

	// Calculate padding based on the longest line (maxLineWidth)
	globalPadding := (m.width - maxLineWidth) / 2
	if globalPadding < 0 {
		globalPadding = 0
	}

	// Now apply colors to the plain ASCII art with animation offset
	coloredAsciiArt := applyAnimatedRainbowColors(plainAsciiArt, m.colorOffset)
	coloredLines := strings.Split(coloredAsciiArt, "\n")

	for _, line := range coloredLines {
		if line == "" {
			continue
		}
		// Use the same padding for all lines
		centeredLines = append(centeredLines, strings.Repeat(" ", globalPadding)+line)
	}

	// Center vertically
	content := strings.Join(centeredLines, "\n")
	verticalPadding := (m.height - len(centeredLines)) / 2
	if verticalPadding < 0 {
		verticalPadding = 0
	}

	result := strings.Repeat("\n", verticalPadding) + content

	// Add quit instruction
	quitPadding := (m.width - 20) / 2
	if quitPadding < 0 {
		quitPadding = 0
	}
	result += "\n\n" + strings.Repeat(" ", quitPadding) + "Press 'q' to quit"

	return result
}

var rootCmd = &cobra.Command{
	Use:   "ccusage-rainbow [text]",
	Short: "Display text in big ASCII art fullscreen",
	Long: `ccusage-rainbow is a CLI tool that displays text in big ASCII art format in fullscreen.
You can provide text as an argument or it will display "HELLO" by default.
Press 'q' to quit the fullscreen view.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := "HELLO"
		if len(args) > 0 {
			text = strings.ToUpper(args[0])
		}

		m := model{text: text}
		p := tea.NewProgram(m, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running program: %v", err)
			os.Exit(1)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func applyAnimatedRainbowColors(text string, offset int) string {
	rainbowColors := []string{
		"#FF0000", // Red
		"#FF8000", // Orange
		"#FFFF00", // Yellow
		"#00FF00", // Green
		"#0080FF", // Blue
		"#4000FF", // Indigo
		"#8000FF", // Violet
	}

	var result strings.Builder
	colorIndex := offset

	for _, char := range text {
		if char == ' ' || char == '\n' {
			result.WriteRune(char)
		} else {
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(rainbowColors[colorIndex%len(rainbowColors)]))
			result.WriteString(style.Render(string(char)))
			colorIndex = (colorIndex + 1) % len(rainbowColors)
		}
	}

	return result.String()
}

func generatePlainBigText(text string) string {
	bigLetters := map[rune][]string{
		'A': {
			" █████ ",
			"██   ██",
			"███████",
			"██   ██",
			"██   ██",
		},
		'B': {
			"██████ ",
			"██   ██",
			"██████ ",
			"██   ██",
			"██████ ",
		},
		'C': {
			" ██████",
			"██     ",
			"██     ",
			"██     ",
			" ██████",
		},
		'D': {
			"██████ ",
			"██   ██",
			"██   ██",
			"██   ██",
			"██████ ",
		},
		'E': {
			"███████",
			"██     ",
			"█████  ",
			"██     ",
			"███████",
		},
		'F': {
			"███████",
			"██     ",
			"█████  ",
			"██     ",
			"██     ",
		},
		'G': {
			" ██████",
			"██     ",
			"██  ███",
			"██   ██",
			" ██████",
		},
		'H': {
			"██   ██",
			"██   ██",
			"███████",
			"██   ██",
			"██   ██",
		},
		'I': {
			"███████",
			"   ██  ",
			"   ██  ",
			"   ██  ",
			"███████",
		},
		'J': {
			"███████",
			"     ██",
			"     ██",
			"██   ██",
			" ██████",
		},
		'K': {
			"██   ██",
			"██  ██ ",
			"█████  ",
			"██  ██ ",
			"██   ██",
		},
		'L': {
			"██     ",
			"██     ",
			"██     ",
			"██     ",
			"███████",
		},
		'M': {
			"███████",
			"██ █ ██",
			"██ █ ██",
			"██   ██",
			"██   ██",
		},
		'N': {
			"███  ██",
			"████ ██",
			"██ ████",
			"██  ███",
			"██   ██",
		},
		'O': {
			" ██████",
			"██    ██",
			"██    ██",
			"██    ██",
			" ██████",
		},
		'P': {
			"██████ ",
			"██   ██",
			"██████ ",
			"██     ",
			"██     ",
		},
		'Q': {
			" ██████",
			"██    ██",
			"██ ██ ██",
			"██  ████",
			" ███████",
		},
		'R': {
			"██████ ",
			"██   ██",
			"██████ ",
			"██   ██",
			"██   ██",
		},
		'S': {
			" ██████",
			"██     ",
			" ██████",
			"     ██",
			" ██████",
		},
		'T': {
			"███████",
			"   ██  ",
			"   ██  ",
			"   ██  ",
			"   ██  ",
		},
		'U': {
			"██   ██",
			"██   ██",
			"██   ██",
			"██   ██",
			" ██████",
		},
		'V': {
			"██   ██",
			"██   ██",
			"██   ██",
			" ██ ██ ",
			"  ███  ",
		},
		'W': {
			"██   ██",
			"██   ██",
			"██ █ ██",
			"██ █ ██",
			"███████",
		},
		'X': {
			"██   ██",
			" ██ ██ ",
			"  ███  ",
			" ██ ██ ",
			"██   ██",
		},
		'Y': {
			"██   ██",
			" ██ ██ ",
			"  ███  ",
			"   ██  ",
			"   ██  ",
		},
		'Z': {
			"███████",
			"    ██ ",
			"   ██  ",
			"  ██   ",
			"███████",
		},
		' ': {
			"       ",
			"       ",
			"       ",
			"       ",
			"       ",
		},
	}

	var result []string
	for i := 0; i < 5; i++ {
		var line string
		for j, char := range text {
			if patterns, ok := bigLetters[char]; ok {
				line += patterns[i]
			} else {
				line += "       "
			}
			// Only add space between characters, not after the last one
			if j < len(text)-1 {
				line += " "
			}
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}
