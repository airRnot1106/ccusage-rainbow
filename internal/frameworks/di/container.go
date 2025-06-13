package di

import (
	"ccusage-rainbow/internal/infrastructure/ascii"
	"ccusage-rainbow/internal/infrastructure/color"
	costInfra "ccusage-rainbow/internal/infrastructure/cost"
	"ccusage-rainbow/internal/interfaces/cli"
	costUseCase "ccusage-rainbow/internal/usecase/cost"
	"ccusage-rainbow/internal/usecase/rainbow"
)

// Container holds all dependencies
type Container struct {
	cliController *cli.Controller
}

// NewContainer creates a new dependency injection container
func NewContainer() *Container {
	// Infrastructure layer
	asciiRenderer := ascii.NewRenderer()
	colorAnimator := color.NewAnimator()
	costService := costInfra.NewService()

	// Use case layer
	rainbowUseCase := rainbow.NewRainbowTextUseCase(asciiRenderer, colorAnimator)
	costDisplayUseCase := costUseCase.NewCostDisplayUseCase(costService)

	// Interface adapters layer
	cliController := cli.NewController(rainbowUseCase, costDisplayUseCase)

	return &Container{
		cliController: cliController,
	}
}

// GetCLIController returns the CLI controller
func (c *Container) GetCLIController() *cli.Controller {
	return c.cliController
}
