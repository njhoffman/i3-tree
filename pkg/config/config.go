package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the i3-tree configuration
type Config struct {
	// DefaultOutputType specifies the default output type: "raw", "all", or "focused"
	DefaultOutputType string `json:"default_output_type"`

	// Display options
	Display DisplayOptions `json:"display"`

	// Formatting options for different node types
	Formatting FormattingOptions `json:"formatting"`

	// Icons for status indicators
	Icons IconOptions `json:"icons"`
}

// DisplayOptions controls what information is shown
type DisplayOptions struct {
	ShowWindowTitles bool     `json:"show_window_titles"`
	ShowMarks        bool     `json:"show_marks"`
	ShowWindowClass  bool     `json:"show_window_class"`
	ShowIcons        bool     `json:"show_icons"`
	Branches         Branches `json:"branches"`
}

// Branches defines the characters used for tree visualization
type Branches struct {
	Horizontal string `json:"horizontal"` // ──
	Vertical   string `json:"vertical"`   // │
	ConnectH   string `json:"connect_h"`  // ├
	ConnectV   string `json:"connect_v"`  // └─
}

// FormattingOptions contains formatting for all node types and layouts
type FormattingOptions struct {
	// Node type formatting
	Root      NodeFormat `json:"root"`
	Output    NodeFormat `json:"output"`
	Workspace NodeFormat `json:"workspace"`
	Con       NodeFormat `json:"con"`
	FloatCon  NodeFormat `json:"float_con"`

	// Layout formatting (consolidated)
	WindowLayout NodeFormat `json:"window_layout"`

	// Window element formatting
	WindowMarks NodeFormat `json:"window_marks"`
	WindowClass NodeFormat `json:"window_class"`
	WindowTitle NodeFormat `json:"window_title"`

	// Focus-specific formatting
	FocusType     NodeFormat `json:"focus_type"`
	FocusBrackets NodeFormat `json:"focus_brackets"`
	FocusBranches NodeFormat `json:"focus_branches"`
	FocusClass    NodeFormat `json:"focus_class"`

	// General elements
	Brackets     NodeFormat `json:"brackets"`
	TreeBranches NodeFormat `json:"tree_branches"`
	Default      NodeFormat `json:"default"`
}

// IconOptions contains icons and their colors for status indicators
type IconOptions struct {
	Fullscreen IconConfig `json:"fullscreen"`
	Floating   IconConfig `json:"floating"`
	Sticky     IconConfig `json:"sticky"`
	Urgent     IconConfig `json:"urgent"`
}

// IconConfig specifies an icon and its formatting
type IconConfig struct {
	Enabled    bool       `json:"enabled"`
	Icon       string     `json:"icon"`
	Foreground int        `json:"foreground"`
	Background int        `json:"background"`
	Attributes Attributes `json:"attributes"`
}

// NodeFormat specifies how a node type or element should be formatted
type NodeFormat struct {
	Foreground int        `json:"foreground"`
	Background int        `json:"background"`
	Attributes Attributes `json:"attributes"`
}

// Attributes contains text formatting attributes
type Attributes struct {
	Bold      bool `json:"bold"`
	Italic    bool `json:"italic"`
	Underline bool `json:"underline"`
	Dim       bool `json:"dim"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		DefaultOutputType: "focused",
		Display: DisplayOptions{
			ShowWindowTitles: true,
			ShowMarks:        true,
			ShowWindowClass:  true,
			ShowIcons:        true,
			Branches: Branches{
				Horizontal: "──",
				Vertical:   "│",
				ConnectH:   "├──",
				ConnectV:   "└──",
			},
		},
		Formatting: FormattingOptions{
			Root: NodeFormat{
				Foreground: 0,
				Background: 0,
				Attributes: Attributes{},
			},
			Output: NodeFormat{
				Foreground: 5,   // magenta (code 35)
				Background: 0,
				Attributes: Attributes{},
			},
			Workspace: NodeFormat{
				Foreground: 6,   // cyan (code 36)
				Background: 0,
				Attributes: Attributes{},
			},
			Con: NodeFormat{
				Foreground: 4,   // blue (code 34)
				Background: 0,
				Attributes: Attributes{},
			},
			FloatCon: NodeFormat{
				Foreground: 4,   // blue (code 34, same as con)
				Background: 0,
				Attributes: Attributes{},
			},
			// Consolidated layout formatting (used for all layouts)
			WindowLayout: NodeFormat{
				Foreground: 3,   // yellow (default, can be overridden)
				Background: 0,
				Attributes: Attributes{},
			},
			// Window element formatting
			WindowMarks: NodeFormat{
				Foreground: 1,   // red (code 31)
				Background: 0,
				Attributes: Attributes{},
			},
			WindowClass: NodeFormat{
				Foreground: 0,   // default
				Background: 0,
				Attributes: Attributes{},
			},
			WindowTitle: NodeFormat{
				Foreground: 0,   // default
				Background: 0,
				Attributes: Attributes{},
			},
			// Focus-specific formatting
			FocusType: NodeFormat{
				Foreground: 0,   // default (will use node type color + bold)
				Background: 0,
				Attributes: Attributes{Bold: true},
			},
			FocusBrackets: NodeFormat{
				Foreground: 0,   // default
				Background: 0,
				Attributes: Attributes{Bold: true},
			},
			FocusBranches: NodeFormat{
				Foreground: 81,  // bright cyan (as requested)
				Background: 0,
				Attributes: Attributes{Bold: true},
			},
			FocusClass: NodeFormat{
				Foreground: 255, // bright white (as requested)
				Background: 0,
				Attributes: Attributes{Bold: true},
			},
			// General elements
			Brackets: NodeFormat{
				Foreground: 0,
				Background: 0,
				Attributes: Attributes{},
			},
			TreeBranches: NodeFormat{
				Foreground: 0,
				Background: 0,
				Attributes: Attributes{},
			},
			Default: NodeFormat{
				Foreground: 0,   // default terminal color
				Background: 0,
				Attributes: Attributes{},
			},
		},
		Icons: IconOptions{
			Fullscreen: IconConfig{
				Enabled:    true,
				Icon:       "󰊓",
				Foreground: 15,  // bright white
				Background: 0,
				Attributes: Attributes{Bold: true},
			},
			Floating: IconConfig{
				Enabled:    true,
				Icon:       "󰭽",
				Foreground: 15,  // bright white
				Background: 0,
				Attributes: Attributes{Bold: true},
			},
			Sticky: IconConfig{
				Enabled:    true,
				Icon:       "󱍭",
				Foreground: 15,  // bright white
				Background: 0,
				Attributes: Attributes{Bold: true},
			},
			Urgent: IconConfig{
				Enabled:    true,
				Icon:       "",
				Foreground: 15,  // bright white
				Background: 0,
				Attributes: Attributes{Bold: true},
			},
		},
	}
}

// Load attempts to load configuration from default paths
// Returns default config if no file is found
func Load() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	// Try default paths in order
	paths := []string{
		filepath.Join(home, ".config", "i3-tree.json"),
		filepath.Join(home, ".config", "i3-tree", "i3-tree.json"),
	}

	for _, path := range paths {
		if config, err := loadFromFile(path); err == nil {
			return config, nil
		}
	}

	// No config file found, create default
	defaultPath := filepath.Join(home, ".config", "i3-tree", "i3-tree.json")
	config := DefaultConfig()

	if err := config.SaveTo(defaultPath); err != nil {
		// If we can't save, just return the default config
		// Don't fail the entire program
		return config, nil
	}

	return config, nil
}

// loadFromFile loads configuration from a specific file
func loadFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", path, err)
	}

	return config, nil
}

// SaveTo saves the configuration to a file
func (c *Config) SaveTo(path string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Marshal with pretty printing
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write to file
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
