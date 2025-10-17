package render

import (
	"fmt"

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
		switch layout {
		case "stacked":
			return au.BrightGreen(s).String()
		case "tabbed":
			return au.Green(s).String()
		case "splith":
			return au.BrightYellow(s).String()
		case "splitv":
			return au.Yellow(s).String()
		default:
			return s
		}
	}

	s := ""

	// only show layout if it has children
	if len(node.Nodes) > 0 {
		return t.wrapBrackets(formatFn(node.Layout, au), isFocused)
	}

	return s
}

func (t *console) formatType(node *i3.Node, au aurora.Aurora, isFocused bool) string {
	if node == nil {
		return ""
	}

	formatFn := func(t i3.NodeType, au aurora.Aurora, bold bool) string {
		s := string(t)

		var colored string
		switch t {
		case "workspace":
			colored = au.Cyan(s).String()
		case "con":
			colored = au.Blue(s).String()
		case "output":
			colored = au.Magenta(s).String()
		default:
			colored = s
		}

		// Make the type text bold if focused
		if bold {
			// We need to apply bold to the already colored text
			// Aurora chaining: color first, then bold
			switch t {
			case "workspace":
				return au.Bold(au.Cyan(s)).String()
			case "con":
				return au.Bold(au.Blue(s)).String()
			case "output":
				return au.Bold(au.Magenta(s)).String()
			default:
				return au.Bold(s).String()
			}
		}

		return colored
	}

	return t.wrapBrackets(formatFn(node.Type, au, isFocused), isFocused)
}
