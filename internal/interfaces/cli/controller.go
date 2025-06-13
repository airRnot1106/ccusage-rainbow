package cli

import (
	"ccusage-rainbow/internal/domain/entities"
	"ccusage-rainbow/internal/interfaces/tui"
	costUseCase "ccusage-rainbow/internal/usecase/cost"
	"ccusage-rainbow/internal/usecase/rainbow"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// Controller handles CLI command execution
type Controller struct {
	rainbowUseCase *rainbow.RainbowTextUseCase
	costUseCase    *costUseCase.CostDisplayUseCase
}

// NewController creates a new CLI controller
func NewController(rainbowUseCase *rainbow.RainbowTextUseCase, costUseCase *costUseCase.CostDisplayUseCase) *Controller {
	return &Controller{
		rainbowUseCase: rainbowUseCase,
		costUseCase:    costUseCase,
	}
}

// CreateRootCommand creates the root cobra command
func (c *Controller) CreateRootCommand() *cobra.Command {
	var useCustomText bool
	var textContent string

	rootCmd := &cobra.Command{
		Use:   "ccusage-rainbow",
		Short: "Display rainbow colored total cost from ccusage",
		Long:  "A CLI tool that fetches total cost from ccusage and displays it as large ASCII text with animated rainbow colors",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runTUI(useCustomText, textContent)
		},
	}

	rootCmd.Flags().BoolVarP(&useCustomText, "custom", "c", false, "Use custom text instead of fetching cost data")
	rootCmd.Flags().StringVarP(&textContent, "text", "t", "HELLO", "Custom text to display (only used with --custom flag)")

	return rootCmd
}

// runTUI starts the TUI application
func (c *Controller) runTUI(useCustomText bool, textContent string) error {
	var text *entities.Text
	var err error

	if useCustomText {
		// Use custom text provided by user
		text = entities.NewText(textContent)
	} else {
		// Fetch cost data and format it
		text, err = c.costUseCase.GetCostText()
		if err != nil {
			// Fallback to error display
			text = entities.NewText("ERROR")
		}
	}

	model := tui.NewModel(text, c.rainbowUseCase)
	program := tea.NewProgram(model, tea.WithAltScreen())

	_, err = program.Run()
	return err
}
