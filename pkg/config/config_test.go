package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/njhoffman/i3-tree/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.DefaultConfig()

	assert.NotNil(t, cfg)
	assert.Equal(t, "focused", cfg.DefaultOutputType)
	assert.True(t, cfg.Display.ShowWindowTitles)
	assert.True(t, cfg.Display.ShowMarks)
	assert.True(t, cfg.Display.ShowWindowClass)
	assert.True(t, cfg.Display.ShowIcons)

	// Check some color defaults
	assert.Equal(t, 6, cfg.Formatting.Workspace.Foreground) // cyan
	assert.Equal(t, 4, cfg.Formatting.Con.Foreground)       // blue
	assert.Equal(t, 1, cfg.Formatting.Marks.Foreground)     // red

	// Check icon defaults
	assert.True(t, cfg.Icons.Fullscreen.Enabled)
	assert.Equal(t, "󰊓", cfg.Icons.Fullscreen.Icon)
	assert.Equal(t, 15, cfg.Icons.Fullscreen.Foreground) // bright white
	assert.True(t, cfg.Icons.Fullscreen.Attributes.Bold)

	assert.True(t, cfg.Icons.Floating.Enabled)
	assert.Equal(t, "󰭽", cfg.Icons.Floating.Icon)
}

func TestConfigSaveAndLoad(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test-config.json")

	// Create a config with custom values
	cfg := config.DefaultConfig()
	cfg.DefaultOutputType = "all"
	cfg.Display.ShowMarks = false
	cfg.Icons.Floating.Icon = "🎈"

	// Save it
	err := cfg.SaveTo(configPath)
	require.NoError(t, err)

	// Verify file exists
	_, err = os.Stat(configPath)
	require.NoError(t, err)

	// Load it back
	data, err := os.ReadFile(configPath)
	require.NoError(t, err)

	// Verify it's valid JSON and contains our custom values
	assert.Contains(t, string(data), "\"default_output_type\": \"all\"")
	assert.Contains(t, string(data), "\"show_marks\": false")
	assert.Contains(t, string(data), "🎈")
}

func TestConfigSaveCreatesDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "nested", "dir", "config.json")

	// Save config to a path that doesn't exist yet
	cfg := config.DefaultConfig()
	err := cfg.SaveTo(configPath)
	require.NoError(t, err)

	// Verify directory was created
	_, err = os.Stat(filepath.Dir(configPath))
	require.NoError(t, err)

	// Verify file exists
	_, err = os.Stat(configPath)
	require.NoError(t, err)
}

func TestLoadNonExistentConfig(t *testing.T) {
	// Save current HOME
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	// Set HOME to a temp directory
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)

	// Load should create default config
	cfg, err := config.Load()
	require.NoError(t, err)
	assert.NotNil(t, cfg)

	// Should have created the default config file
	defaultPath := filepath.Join(tmpDir, ".config", "i3-tree", "i3-tree.json")
	_, err = os.Stat(defaultPath)
	assert.NoError(t, err)
}
