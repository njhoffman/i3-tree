# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
