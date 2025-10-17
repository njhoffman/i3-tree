package internal

import (
	"os"

	"github.com/eh-am/i3-tree/pkg/config"
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/eh-am/i3-tree/pkg/render"
)

type RendererStrat string

var (
	// Console strategy
	ConsoleStrat RendererStrat = "console"
	// Console, but no color strategy
	ConsoleNoColorStrat RendererStrat = "no-color"

	// List of all available render strategies
	AvailableRendererStrats = []RendererStrat{
		ConsoleStrat,
		ConsoleNoColorStrat,
	}
)

// NewRenderer creates a i3treeviewer.Renderer
// Based on a strategy
// Otherwise it fails with BadStratError
func NewRenderer(strat string) (i3treeviewer.Renderer, error) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		// If config loading fails, use default config
		cfg = config.DefaultConfig()
	}

	switch RendererStrat(strat) {
	case ConsoleStrat:
		return render.NewColoredConsoleWithConfig(os.Stdout, cfg), nil

	case ConsoleNoColorStrat:
		return render.NewMonochromaticConsoleWithConfig(os.Stdout, cfg), nil

	default:
		return nil, BadStratError{strat}
	}
}
