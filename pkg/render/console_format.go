package render

import (
	"fmt"

	"github.com/njhoffman/i3-tree/pkg/config"
	"github.com/logrusorgru/aurora"
	"go.i3wm.org/i3/v4"
)

func (t *console) wrapBrackets(s string, bold bool) string {
	if s != "" {
		if bold {
			return fmt.Sprint(t.au.Bold("[").String(), s, t.au.Bold("]").String())
		}
		return fmt.Sprint("[", s, "]")
	}
	return s
}

func (t *console) formatLayout(node *i3.Node, au aurora.Aurora, isFocused bool) string {
	if node == nil {
		return ""
	}

	formatFn := func(layout i3.Layout, au aurora.Aurora) string {
		s := string(layout)
		// Use consolidated window_layout formatting for all layouts
		return t.config.Formatting.WindowLayout.ApplyFormat(s, au)
	}

	s := ""

	// only show layout if it has children
	if len(node.Nodes) > 0 {
		return t.wrapBrackets(formatFn(node.Layout, au), isFocused)
	}

	return s
}

func (t *console) formatType(node *i3.Node, au aurora.Aurora, isFocused bool, isFloating bool) string {
	if node == nil {
		return ""
	}

	formatFn := func(nodeType i3.NodeType, au aurora.Aurora, bold bool, floating bool) string {
		s := string(nodeType)

		// Replace "con" with "fcon" for floating containers
		// Also replace "floating_con" with "fcon"
		if nodeType == "floating_con" || (floating && nodeType == "con") {
			s = "fcon"
		}

		// Use config formatting for node types
		var nodeFormat *config.NodeFormat
		switch nodeType {
		case "workspace":
			nodeFormat = &t.config.Formatting.Workspace
		case "con":
			nodeFormat = &t.config.Formatting.Con
		case "floating_con":
			nodeFormat = &t.config.Formatting.FloatCon
		case "output":
			nodeFormat = &t.config.Formatting.Output
		case "root":
			nodeFormat = &t.config.Formatting.Root
		default:
			return s
		}

		if nodeFormat != nil {
			// If focused, we need to apply bold to the formatting
			if bold {
				// Create a copy with bold attribute
				boldFormat := *nodeFormat
				boldFormat.Attributes.Bold = true
				return boldFormat.ApplyFormat(s, au)
			}
			return nodeFormat.ApplyFormat(s, au)
		}

		return s
	}

	return t.wrapBrackets(formatFn(node.Type, au, isFocused, isFloating), isFocused)
}
