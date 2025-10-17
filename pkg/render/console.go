package render

import (
	"fmt"
	"io"
	"strings"

	"github.com/eh-am/i3-tree/pkg/config"
	"github.com/logrusorgru/aurora"
	"go.i3wm.org/i3/v4"
)

type ColoredConsole struct {
	*console
}
type MonochromaticConsole struct {
	*console
}

type console struct {
	w      io.Writer
	au     aurora.Aurora
	config *config.Config
}

func NewColoredConsole(w io.Writer) ColoredConsole {
	return NewColoredConsoleWithConfig(w, config.DefaultConfig())
}

func NewColoredConsoleWithConfig(w io.Writer, cfg *config.Config) ColoredConsole {
	return ColoredConsole{
		newConsole(w, true, cfg),
	}
}

func NewMonochromaticConsole(w io.Writer) MonochromaticConsole {
	return NewMonochromaticConsoleWithConfig(w, config.DefaultConfig())
}

func NewMonochromaticConsoleWithConfig(w io.Writer, cfg *config.Config) MonochromaticConsole {
	return MonochromaticConsole{
		newConsole(w, false, cfg),
	}
}

func newConsole(w io.Writer, colors bool, cfg *config.Config) *console {
	return &console{
		w:      w,
		au:     aurora.NewAurora(colors),
		config: cfg,
	}
}

func (t *console) Render(tree *i3.Tree) {
	// Build a set of node IDs that are on the path to the focused node
	focusedPath := t.buildFocusedPath(tree.Root)
	t.print(tree.Root, "", "", 0, focusedPath, false, false)
}

// buildFocusedPath finds the path from root to the focused node
// Returns a map of node IDs that are on this path
func (t *console) buildFocusedPath(node *i3.Node) map[i3.NodeID]bool {
	path := make(map[i3.NodeID]bool)
	t.findFocusedPath(node, path)
	return path
}

// findFocusedPath recursively searches for the focused node
// and marks all nodes on the path to it
// Returns true if this node or any of its children is/contains the focused node
func (t *console) findFocusedPath(node *i3.Node, path map[i3.NodeID]bool) bool {
	if node == nil {
		return false
	}

	// Check if this node is focused
	if node.Focused {
		path[node.ID] = true
		return true
	}

	// Check if any child contains the focused node (regular nodes)
	for _, child := range node.Nodes {
		if t.findFocusedPath(child, path) {
			// This node is on the path to the focused node
			path[node.ID] = true
			return true
		}
	}

	// Check floating nodes too
	for _, child := range node.FloatingNodes {
		if t.findFocusedPath(child, path) {
			// This node is on the path to the focused node
			path[node.ID] = true
			return true
		}
	}

	return false
}

// formatWindowDetails formats additional window information like icons, class, marks, and title
// Icons are displayed first, followed by class, title, and marks
func (t *console) formatWindowDetails(node *i3.Node, isFloating bool) string {
	if node == nil {
		return ""
	}

	var result string

	// Build status icons first if enabled
	icons := ""
	if t.config.Display.ShowIcons {
		// Fullscreen icon
		if t.config.Icons.Fullscreen.Enabled && node.FullscreenMode != 0 {
			icon := t.config.Icons.Fullscreen.ApplyFormat(t.config.Icons.Fullscreen.Icon, t.au)
			icons += " " + icon
		}

		// Floating icon
		if t.config.Icons.Floating.Enabled && (isFloating || node.Type == "floating_con") {
			icon := t.config.Icons.Floating.ApplyFormat(t.config.Icons.Floating.Icon, t.au)
			icons += " " + icon
		}

		// Sticky icon - Note: i3 doesn't expose sticky directly in the tree
		// We check if a window has the special mark "_sticky" which is often used
		isSticky := false
		for _, mark := range node.Marks {
			if mark == "_sticky" {
				isSticky = true
				break
			}
		}
		if t.config.Icons.Sticky.Enabled && isSticky {
			icon := t.config.Icons.Sticky.ApplyFormat(t.config.Icons.Sticky.Icon, t.au)
			icons += " " + icon
		}

		// Urgent icon
		if t.config.Icons.Urgent.Enabled && node.Urgent {
			icon := t.config.Icons.Urgent.ApplyFormat(t.config.Icons.Urgent.Icon, t.au)
			icons += " " + icon
		}
	}

	// Add icons first
	result += icons

	// Add window class if available (only for con type)
	if t.config.Display.ShowWindowClass && node.Type == "con" && node.WindowProperties.Class != "" {
		className := fmt.Sprintf("(%s)", node.WindowProperties.Class)
		// Apply focus_class formatting if this is a focused node
		if node.Focused {
			className = t.config.Formatting.FocusClass.ApplyFormat(className, t.au)
		} else {
			className = t.config.Formatting.WindowClass.ApplyFormat(className, t.au)
		}
		result += " " + className
	}

	// Add window title
	if t.config.Display.ShowWindowTitles && node.Name != "" {
		// Truncate title if too long (> 80 chars)
		title := node.Name
		maxLen := 80
		if len(title) > maxLen {
			title = title[:maxLen-3] + "..."
		}
		formattedTitle := t.config.Formatting.WindowTitle.ApplyFormat(title, t.au)
		result += " " + formattedTitle
	}

	// Add marks in configured color brackets
	if t.config.Display.ShowMarks && len(node.Marks) > 0 {
		marksStr := ""
		for i, mark := range node.Marks {
			if i > 0 {
				marksStr += ", "
			}
			marksStr += mark
		}
		formattedMarks := t.config.Formatting.WindowMarks.ApplyFormat(fmt.Sprintf("[%s]", marksStr), t.au)
		result += " " + formattedMarks
	}

	return result
}

func (t *console) print(node *i3.Node, prefix string, marker string, level int, focusedPath map[i3.NodeID]bool, isFloating bool, hasFocusedSibling bool) {
	if node == nil {
		return
	}

	isOnFocusedPath := focusedPath[node.ID]
	isFocused := node.Focused

	// Special handling for floating_con: collapse it with its child
	if node.Type == "floating_con" && len(node.Nodes) == 1 {
		child := node.Nodes[0]

		// Apply focus_branches formatting
		displayMarker := marker
		if marker != "" {
			if isOnFocusedPath {
				// On the actual focused path: highlight entire marker
				displayMarker = t.config.Formatting.FocusBranches.ApplyFormat(marker, t.au)
			} else if hasFocusedSibling {
				// Has focused sibling: highlight only connector
				connector := marker[:len(t.config.Display.Branches.ConnectH)]
				horizontal := marker[len(connector):]
				formattedConnector := t.config.Formatting.FocusBranches.ApplyFormat(connector, t.au)
				displayMarker = formattedConnector + horizontal
			}
		}

		// Format the type as fcon
		ftype := t.formatType(node, t.au, child.Focused, true)

		// Get child's window details (which will include icons first)
		windowDetails := t.formatWindowDetails(child, true)

		fmt.Fprint(
			t.w,
			prefix,
			displayMarker,
			ftype,
			windowDetails,
			"\n",
		)
		return
	}

	ftype := t.formatType(node, t.au, isFocused, isFloating)
	flayout := t.formatLayout(node, t.au, isFocused)

	// Apply focus_branches formatting to marker
	displayMarker := marker
	if marker != "" {
		if isOnFocusedPath {
			// On the actual focused path: highlight entire marker (connector + horizontal)
			displayMarker = t.config.Formatting.FocusBranches.ApplyFormat(marker, t.au)
		} else if hasFocusedSibling {
			// Has focused sibling (path continuation): highlight only connector
			connector := marker[:len(t.config.Display.Branches.ConnectH)]
			horizontal := marker[len(connector):]
			formattedConnector := t.config.Formatting.FocusBranches.ApplyFormat(connector, t.au)
			displayMarker = formattedConnector + horizontal
		}
	}

	// Format additional window details (class, marks, icons)
	windowDetails := t.formatWindowDetails(node, isFloating)

	fmt.Fprint(
		t.w,
		prefix,
		displayMarker,
		ftype,
		flayout,
		windowDetails,
		"\n",
	)

	// Combine regular nodes and floating nodes
	allNodes := append([]*i3.Node{}, node.Nodes...)
	allNodes = append(allNodes, node.FloatingNodes...)

	for i, n := range allNodes {
		newPrefix := prefix
		newMarker := ""

		// Check if this is a floating node
		childIsFloating := i >= len(node.Nodes)

		// Build the marker: connector + horizontal
		var connector string
		if i == len(allNodes)-1 {
			connector = t.config.Display.Branches.ConnectV // └
		} else {
			connector = t.config.Display.Branches.ConnectH // ├
		}
		newMarker = connector + t.config.Display.Branches.Horizontal // e.g., "├──"

		// Check if any subsequent sibling is on the focused path
		// This determines if the trunk should be highlighted
		anySiblingOnFocusedPath := false
		for j := i + 1; j < len(allNodes); j++ {
			if focusedPath[allNodes[j].ID] {
				anySiblingOnFocusedPath = true
				break
			}
		}

		// Determine prefix for child based on current marker
		// Check if current marker starts with ConnectH (we're a middle node)
		// Use strings.HasPrefix to properly handle UTF-8 characters
		if strings.HasPrefix(marker, t.config.Display.Branches.ConnectH) {
			// This node has more children, so add trunk to prefix
			// Highlight the trunk if this node has a focused sibling (passed from parent)
			var trunkChar string
			if hasFocusedSibling {
				formattedVertical := t.config.Formatting.FocusBranches.ApplyFormat(t.config.Display.Branches.Vertical, t.au)
				trunkChar = formattedVertical + "  "
			} else {
				trunkChar = t.config.Display.Branches.Vertical + "  "
			}
			newPrefix = newPrefix + trunkChar
		} else {
			// don't indent starting from root
			if level == 0 {
				newPrefix = ""
			} else {
				newPrefix = newPrefix + "   "
			}
		}

		t.print(n, newPrefix, newMarker, level+1, focusedPath, childIsFloating, anySiblingOnFocusedPath)
	}
}
