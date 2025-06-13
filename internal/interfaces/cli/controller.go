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
	var useBankruptMode bool
	var useHiMode bool

	rootCmd := &cobra.Command{
		Use:   "ccusage-rainbow",
		Short: "Display rainbow colored total cost from ccusage",
		Long:  "A CLI tool that fetches total cost from ccusage and displays it as large ASCII text with animated rainbow colors",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runTUI(useBankruptMode, useHiMode)
		},
	}

	rootCmd.Flags().BoolVarP(&useBankruptMode, "bankrupt", "", false, "")
	_ = rootCmd.Flags().MarkHidden("bankrupt")
	rootCmd.Flags().BoolVarP(&useHiMode, "hi", "", false, "")
	_ = rootCmd.Flags().MarkHidden("hi")

	return rootCmd
}

// runTUI starts the TUI application
func (c *Controller) runTUI(useBankruptMode bool, useHiMode bool) error {
	var text *entities.Text
	var err error

	if useHiMode {
		// Hidden option to display "HELLO"
		text = entities.NewText("HELLO")
	} else if useBankruptMode {
		// Hidden bankrupt mode - display large cost
		text = entities.NewText("$9999.99")
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
