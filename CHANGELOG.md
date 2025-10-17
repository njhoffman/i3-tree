# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
