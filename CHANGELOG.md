# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2025-10-16

### Added
- Configurable tree branch characters (horizontal, vertical, connect_h, connect_v)
- New formatting options for window elements: `window_class`, `window_title`
- Focus-specific formatting options: `focus_type`, `focus_brackets`, `focus_branches`, `focus_class`
- Default formatting option for all other text

### Changed
- **BREAKING**: Config field `marks` renamed to `window_marks` in formatting options
- **BREAKING**: Individual layout configs (splith, splitv, tabbed, stacked) replaced with single `window_layout` config
- Focus branches now display in bright cyan (color 81) by default
- Focused window class now displays in bright white (color 255) by default
- All layout types now use consistent yellow coloring (color 33) by default

## [1.0.0] - 2025-10-16

### Added
- Watch mode with `--watch` / `-w` flag to continuously refresh the tree display
- Configurable refresh interval (default: 5 seconds)
- Examples: `i3-tree --watch=0` (5 sec default), `i3-tree -w 2` (2 sec interval)

## [0.8.0] - 2025-10-16

### Changed
- Floating windows now display on a single line (collapsed with parent floating_con)
- Status icons now appear before window class instead of after title
- Floating containers no longer show layout information

## [0.7.0] - 2025-10-16

### Added
- Configuration file support with JSON format
- Auto-creation of default config file at `~/.config/i3-tree/i3-tree.json`
- Configurable colors for all node types and layouts (using 0-255 ANSI color codes)
- Configurable text attributes (bold, italic, underline, dim)
- Configurable status icons (fullscreen, floating, sticky, urgent) with customizable colors
- Display options to toggle window titles, marks, window class, and icons
- Default output type configuration ("focused", "all", or "raw")

## [0.6.0] - 2025-10-16

### Added
- Display floating icon (󰭽) in bright white for floating containers

## [0.5.0] - 2025-10-16

### Added
- Display floating windows in the tree with `[fcon]` type instead of `[con]`
- Include FloatingNodes in tree traversal and rendering
- Support for focused path highlighting through floating windows
- Added comprehensive tests for floating windows display

## [0.4.0] - 2025-10-16

### Added
- Display window class name in parentheses after window type (e.g., `[con] (Alacritty)`)
- Display window marks in red brackets after window title (e.g., `[_last, scratch]`)
- Display fullscreen icon (󰊓) in bright white for fullscreen windows
- Display sticky icon (󱍭) in bright white for windows with `_sticky` mark
- Display urgent icon () in bright white for windows with urgent hint
- Automatic title truncation for long window titles (max 80 characters)
- Added comprehensive tests for window details display

## [0.3.0] - 2025-10-16

### Added
- Make focused window type text bold (e.g., the text "con" in `[con]` is now bold for focused windows)
- Updated tests to reflect bold type text in focused windows

## [0.2.0] - 2025-10-16

### Added
- Highlight focused branch: tree branches and brackets are now displayed in bold from root to the currently focused window
- Added tests for focused branch highlighting feature

## [0.1.0] - 2025-10-16

### Added
- Initial versioning and CHANGELOG.md
- Display focused workspace by default
- Support for displaying all non-empty workspaces with `i3-tree all`
- Support for displaying specific workspace by number (e.g., `i3-tree 6`)
- Support for raw tree output with `i3-tree raw`
- Colored console output (default)
- No-color output option with `--render=no-color`
- Mock data support with `--from=mock` for testing without i3
