package render

import (
	"fmt"
	"io"

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
	w  io.Writer
	au aurora.Aurora
}

func NewColoredConsole(w io.Writer) ColoredConsole {
	return ColoredConsole{
		newConsole(w, true),
	}
}

func NewMonochromaticConsole(w io.Writer) MonochromaticConsole {
	return MonochromaticConsole{
		newConsole(w, false),
	}
}

func newConsole(w io.Writer, colors bool) *console {
	return &console{
		w:  w,
		au: aurora.NewAurora(colors),
	}
}

func (t *console) Render(tree *i3.Tree) {
	// Build a set of node IDs that are on the path to the focused node
	focusedPath := t.buildFocusedPath(tree.Root)
	t.print(tree.Root, "", "", 0, focusedPath)
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

	// Check if any child contains the focused node
	for _, child := range node.Nodes {
		if t.findFocusedPath(child, path) {
			// This node is on the path to the focused node
			path[node.ID] = true
			return true
		}
	}

	return false
}

func (t *console) print(node *i3.Node, prefix string, marker string, level int, focusedPath map[i3.NodeID]bool) {
	if node == nil {
		return
	}

	isOnFocusedPath := focusedPath[node.ID]
	isFocused := node.Focused

	ftype := t.formatType(node, t.au, isFocused)
	flayout := t.formatLayout(node, t.au, isFocused)

	// Make the marker bold if on focused path
	displayMarker := marker
	if isOnFocusedPath && marker != "" {
		displayMarker = t.au.Bold(marker).String()
	}

	fmt.Fprint(
		t.w,
		prefix,
		displayMarker,
		ftype,
		flayout,
		" ",
		node.Name,
		"\n",
	)

	for i, n := range node.Nodes {
		newPrefix := prefix
		newMarker := ""

		// figure out what's the marker for the next iteration
		if i == len(node.Nodes)-1 {
			newMarker = "└──" // last node
		} else {
			newMarker = "├──" // middle node
		}

		// Determine the trunk character
		trunkChar := "│  "
		spaceChar := "   "

		// Make trunk bold if this node is on the focused path and the child is too
		childOnFocusedPath := focusedPath[n.ID]
		if isOnFocusedPath && childOnFocusedPath {
			trunkChar = t.au.Bold("│").String() + "  "
		}

		// i am currently a middle node
		if marker == "├──" {
			// so my children should display my trunk
			newPrefix = newPrefix + trunkChar
		} else {
			// don't ident starting from root
			if level == 0 {
				newPrefix = ""
			} else {
				newPrefix = newPrefix + spaceChar
			}
		}

		t.print(n, newPrefix, newMarker, level+1, focusedPath)
	}
}
