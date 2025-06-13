package tui

import (
	"strings"
	"time"

	"ccusage-rainbow/internal/domain/entities"
	"ccusage-rainbow/internal/domain/interfaces"
	"ccusage-rainbow/internal/usecase/rainbow"

	tea "github.com/charmbracelet/bubbletea"
)

// TickMsg represents a timer tick for animation
type TickMsg time.Time

// Model represents the TUI model following Clean Architecture
type Model struct {
	text       *entities.Text
	useCase    *rainbow.RainbowTextUseCase
	dimensions interfaces.DisplayDimensions
}

// NewModel creates a new TUI model
func NewModel(text *entities.Text, useCase *rainbow.RainbowTextUseCase) *Model {
	return &Model{
		text:    text,
		useCase: useCase,
	}
}

// Init initializes the model
func (m *Model) Init() tea.Cmd {
	return tea.Tick(m.useCase.GetAnimationInterval(), func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Update handles messages and updates the model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.dimensions = interfaces.DisplayDimensions{
			Width:  msg.Width,
			Height: msg.Height,
		}
	case TickMsg:
		m.useCase.AdvanceAnimation()
		return m, tea.Tick(m.useCase.GetAnimationInterval(), func(t time.Time) tea.Msg {
			return TickMsg(t)
		})
	}
	return m, nil
}

// View renders the current view
func (m *Model) View() string {
	// Handle case when dimensions are not set yet
	if m.dimensions.Width <= 0 || m.dimensions.Height <= 0 {
		return "Loading..."
	}

	// Select optimal font size based on terminal dimensions
	fontSize, err := m.useCase.SelectOptimalFontSize(m.text, m.dimensions.Width, m.dimensions.Height)
	if err != nil {
		return "Error: " + err.Error()
	}

	// Get display width for centering calculation
	displayWidth, err := m.useCase.GetDisplayWidthWithSize(m.text, fontSize)
	if err != nil {
		return "Error: " + err.Error()
	}

	// If text still doesn't fit even with smallest size, show fallback message
	if displayWidth > m.dimensions.Width {
		fallbackMsg := "Terminal too small"
		padding := (m.dimensions.Width - len(fallbackMsg)) / 2
		if padding < 0 {
			padding = 0
		}
		verticalPadding := m.dimensions.Height / 2
		result := strings.Repeat("\n", verticalPadding) + strings.Repeat(" ", padding) + fallbackMsg
		return result
	}

	// Render animated text with selected font size
	coloredText, err := m.useCase.RenderAnimatedTextWithSize(m.text, fontSize)
	if err != nil {
		return "Error: " + err.Error()
	}

	// Center the text
	lines := strings.Split(coloredText, "\n")
	var centeredLines []string

	// Calculate global padding for consistent alignment
	globalPadding := (m.dimensions.Width - displayWidth) / 2
	if globalPadding < 0 {
		globalPadding = 0
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		centeredLines = append(centeredLines, strings.Repeat(" ", globalPadding)+line)
	}

	// Center vertically
	content := strings.Join(centeredLines, "\n")
	verticalPadding := (m.dimensions.Height - len(centeredLines)) / 2
	if verticalPadding < 0 {
		verticalPadding = 0
	}

	result := strings.Repeat("\n", verticalPadding) + content

	return result
}
