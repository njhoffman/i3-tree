package config

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// ApplyFormat applies a NodeFormat to a string using Aurora
func (nf NodeFormat) ApplyFormat(s string, au aurora.Aurora) string {
	// Start with the base string wrapped in Aurora
	var result aurora.Value = au.Reset(s)

	// Apply foreground color if specified (1-256)
	if nf.Foreground > 0 {
		result = colorize(result, nf.Foreground, false, au)
	}

	// Apply background color if specified (1-256)
	if nf.Background > 0 {
		result = colorize(result, nf.Background, true, au)
	}

	// Apply text attributes
	if nf.Attributes.Bold {
		result = au.Bold(result)
	}
	if nf.Attributes.Italic {
		result = au.Italic(result)
	}
	if nf.Attributes.Underline {
		result = au.Underline(result)
	}
	if nf.Attributes.Dim {
		result = au.Faint(result)
	}

	return result.String()
}

// ApplyIconFormat applies an IconConfig format to a string using Aurora
func (ic IconConfig) ApplyFormat(s string, au aurora.Aurora) string {
	nf := NodeFormat{
		Foreground: ic.Foreground,
		Background: ic.Background,
		Attributes: ic.Attributes,
	}
	return nf.ApplyFormat(s, au)
}

// colorize applies an ANSI color (1-256) to an Aurora value
func colorize(v aurora.Value, color int, background bool, au aurora.Aurora) aurora.Value {
	// Map ANSI colors to Aurora colors
	// Note: Aurora has built-in support for 3-bit/4-bit colors
	// For 8-bit (256) colors, we use Index functions

	if color < 1 || color > 256 {
		return v
	}

	// For standard 16 colors (1-16), map to Aurora's named colors
	if color <= 16 {
		return applyStandardColor(v, color, background, au)
	}

	// For 256 colors (17-256), use Aurora's Index functions
	if background {
		return v.BgIndex(uint8(color - 1)) // Aurora uses 0-255 indexing
	}
	return v.Index(uint8(color - 1))
}

// applyStandardColor applies standard ANSI colors (0-15)
func applyStandardColor(v aurora.Value, color int, background bool, au aurora.Aurora) aurora.Value {
	// Standard ANSI colors (0-7) and bright variants (8-15)
	// Note: ANSI color codes are typically 30-37 (foreground) and 40-47 (background)
	// with bright variants at 90-97 and 100-107
	// Our config uses 0-15 mapping directly to ANSI color offsets

	if background {
		switch color {
		case 0: // black
			return au.BgBlack(v)
		case 1: // red
			return au.BgRed(v)
		case 2: // green
			return au.BgGreen(v)
		case 3: // yellow
			return au.BgYellow(v)
		case 4: // blue
			return au.BgBlue(v)
		case 5: // magenta
			return au.BgMagenta(v)
		case 6: // cyan
			return au.BgCyan(v)
		case 7: // white
			return au.BgWhite(v)
		case 8: // bright black (gray)
			return au.BgBrightBlack(v)
		case 9: // bright red
			return au.BgBrightRed(v)
		case 10: // bright green
			return au.BgBrightGreen(v)
		case 11: // bright yellow
			return au.BgBrightYellow(v)
		case 12: // bright blue
			return au.BgBrightBlue(v)
		case 13: // bright magenta
			return au.BgBrightMagenta(v)
		case 14: // bright cyan
			return au.BgBrightCyan(v)
		case 15: // bright white
			return au.BgBrightWhite(v)
		default:
			return v
		}
	}

	// Foreground colors
	switch color {
	case 0: // black
		return au.Black(v)
	case 1: // red
		return au.Red(v)
	case 2: // green
		return au.Green(v)
	case 3: // yellow
		return au.Yellow(v)
	case 4: // blue
		return au.Blue(v)
	case 5: // magenta
		return au.Magenta(v)
	case 6: // cyan
		return au.Cyan(v)
	case 7: // white
		return au.White(v)
	case 8: // bright black (gray)
		return au.BrightBlack(v)
	case 9: // bright red
		return au.BrightRed(v)
	case 10: // bright green
		return au.BrightGreen(v)
	case 11: // bright yellow
		return au.BrightYellow(v)
	case 12: // bright blue
		return au.BrightBlue(v)
	case 13: // bright magenta
		return au.BrightMagenta(v)
	case 14: // bright cyan
		return au.BrightCyan(v)
	case 15: // bright white
		return au.BrightWhite(v)
	default:
		return v
	}
}

// GetANSICode returns the ANSI escape code for a NodeFormat
// This is useful for debugging or when you need the raw code
func (nf NodeFormat) GetANSICode() string {
	var codes []string

	// Foreground color
	if nf.Foreground > 0 {
		if nf.Foreground <= 16 {
			codes = append(codes, fmt.Sprintf("%d", 29+nf.Foreground))
		} else {
			codes = append(codes, fmt.Sprintf("38;5;%d", nf.Foreground-1))
		}
	}

	// Background color
	if nf.Background > 0 {
		if nf.Background <= 16 {
			codes = append(codes, fmt.Sprintf("%d", 39+nf.Background))
		} else {
			codes = append(codes, fmt.Sprintf("48;5;%d", nf.Background-1))
		}
	}

	// Attributes
	if nf.Attributes.Bold {
		codes = append(codes, "1")
	}
	if nf.Attributes.Dim {
		codes = append(codes, "2")
	}
	if nf.Attributes.Italic {
		codes = append(codes, "3")
	}
	if nf.Attributes.Underline {
		codes = append(codes, "4")
	}

	if len(codes) == 0 {
		return ""
	}

	result := "\x1b["
	for i, code := range codes {
		if i > 0 {
			result += ";"
		}
		result += code
	}
	result += "m"
	return result
}
