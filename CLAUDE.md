# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

i3-tree is a CLI tool that displays the i3 window manager's tree structure in a user-friendly format, similar to the Unix `tree` command. It visualizes workspaces, windows, outputs, and their hierarchical relationships.

## Development Commands

### Testing
```bash
# Run all tests
make test
go test ./...

# Run tests with coverage report
make cover
# This creates tmp/cover.html for viewing coverage

# Run tests for a specific package
go test ./pkg/prune
go test ./cmd/internal
```

### Building and Installing
```bash
# Install dependencies
make install
go get

# Build and install the binary
go install

# Run locally without installing
go run main.go
go run main.go --from=mock  # Use mock data for testing
```

### Linting
```bash
# Run golangci-lint
make lint
golangci-lint run ./...
```

## Architecture

The application follows a strategy pattern with three main pluggable components:

### Core Pipeline (pkg/i3treeviewer/i3treeviewer.go)

The `i3TreeViewer` orchestrates the three-step pipeline:
1. **Fetch**: Retrieve the i3 tree
2. **Prune**: Filter/transform the tree based on user requirements
3. **Render**: Display the tree to the user

Each step is an interface, allowing different strategies to be plugged in.

### Component Interfaces

- **Fetcher** (`pkg/fetch/`): Retrieves i3 tree data
  - `FromI3`: Fetches from live i3 instance via `go.i3wm.org/i3/v4`
  - `FromFake`: Returns mock data for testing without i3

- **Pruner** (`pkg/prune/`): Filters tree to show relevant information
  - `FocusedWs`: Shows only the currently focused workspace (default)
  - `NonEmptyWs`: Shows all non-empty workspaces
  - `Ws`: Shows a specific workspace by name/number
  - `NoOp`: Shows the raw unfiltered tree

- **Renderer** (`pkg/render/`): Outputs the tree
  - `ColoredConsole`: Colored terminal output (default)
  - `MonochromaticConsole`: Plain text without colors

### Strategy Selection (cmd/internal/)

The `cmd/internal` package contains factory functions that map CLI flags and arguments to concrete strategy implementations:
- `fetch_strats.go`: Maps `--from` flag to Fetcher implementations
- `prune_strats.go`: Maps positional arguments to Pruner implementations
- `render_strats.go`: Maps `--render` flag to Renderer implementations

### CLI Entry Point (cmd/root.go)

Uses `peterbourgon/ff/v3` for flag parsing. The `rootExec` function:
1. Creates strategies based on flags/args
2. Instantiates `i3TreeViewer` with the strategies
3. Calls `View()` to execute the pipeline

## Key Design Patterns

- **Strategy Pattern**: Core architecture allows swapping Fetcher, Pruner, and Renderer implementations
- **Factory Pattern**: `NewFetcher()`, `NewPruner()`, `NewRenderer()` create concrete implementations
- **Interface Segregation**: Small, focused interfaces (`Fetcher`, `Pruner`, `Renderer`)

## Testing Strategy

- Strategy implementations have dedicated test files (e.g., `focusedws_test.go`)
- Factory functions are tested in `*_strats_test.go` files
- Mock/fake implementations enable testing without i3 running

## Dependencies

- `go.i3wm.org/i3/v4`: i3 IPC library for fetching window tree
- `github.com/peterbourgon/ff/v3`: CLI flag parsing
- `github.com/logrusorgru/aurora`: Terminal color output
- `github.com/stretchr/testify`: Testing assertions
