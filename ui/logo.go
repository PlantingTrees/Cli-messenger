package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// --- Configuration ---
const (
	// Gradient 1: The "Fire" Center (YE)
	gradFireStart = "#F5F5DC" // White
	gradFireEnd   = "#FF0055" // Fire Purple

	// Gradient 2: The "Royal" Sides ("0", "H", Quotes)
	gradRoyalStart = "#F5F5DC" // Beige (Corrected hex for Beige)
	gradRoyalEnd   = "#6037C8" // Deep Purple
)

// --- The Blocks ---
// (Using the exact shapes you provided)

var blockQuote = []string{
	"██╗██╗",
	"██║██║",
	"╚═╝╚═╝",
	"      ",
	"      ",
	"      ",
}

var blockZero = []string{
	" ██████╗ ",
	"██╔═████╗",
	"██║██╔██║",
	"████╔╝██║",
	"╚██████╔╝",
	" ╚═════╝ ",
}

var blockYE = []string{
	"██╗   ██╗███████╗",
	"╚██╗ ██╔╝██╔════╝",
	" ╚████╔╝ █████╗  ",
	"  ╚██╔╝  ██╔══╝  ",
	"   ██║   ███████╗",
	"   ╚═╝   ╚══════╝",
}

var blockH = []string{
	"██╗  ██╗",
	"██║  ██║",
	"███████║",
	"██╔══██║",
	"██║  ██║",
	"╚═╝  ╚═╝",
}

// --- Main Rendering Function ---

// FIX: Now accepts 'frame' (an integer that grows over time)
func RenderLogo(frame int) string {
	// We pass the frame to the gradient renderer
	quoteLeft := renderAnimatedGradient(blockQuote, gradRoyalStart, gradRoyalEnd, frame)
	quoteRight := renderAnimatedGradient(blockQuote, gradRoyalStart, gradRoyalEnd, frame)
	zero := renderAnimatedGradient(blockZero, gradRoyalStart, gradRoyalEnd, frame)
	h := renderAnimatedGradient(blockH, gradRoyalStart, gradRoyalEnd, frame)

	ye := renderAnimatedGradient(blockYE, gradFireStart, gradFireEnd, frame)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		quoteLeft,
		"  ",
		zero,
		" ",
		ye,
		" ",
		h,
		"  ",
		quoteRight,
	)
}

// --- Animated Gradient Logic ---

func renderAnimatedGradient(lines []string, startColor, endColor string, frame int) string {
	var styledLines []string
	r1, g1, b1 := hexToRGB(startColor)
	r2, g2, b2 := hexToRGB(endColor)

	for i, line := range lines {
		// --- THE MATH MAGIC ---
		// We use a Sine wave to oscillate the color mix between 0.0 and 1.0.
		// float64(i)*0.3  -> Makes each row a slightly different color (Gradient)
		// float64(frame)*0.1 -> Moves the wave over time (Animation Speed)
		t := 0.5 * (1 + math.Sin(float64(i)*0.3-float64(frame)*0.15))

		// Interpolate RGB
		r := uint8(float64(r1) + t*(float64(r2)-float64(r1)))
		g := uint8(float64(g1) + t*(float64(g2)-float64(g1)))
		b := uint8(float64(b1) + t*(float64(b2)-float64(b1)))

		c := lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", r, g, b))
		style := lipgloss.NewStyle().Foreground(c)
		styledLines = append(styledLines, style.Render(line))
	}

	return strings.Join(styledLines, "\n")
}

func hexToRGB(hex string) (uint8, uint8, uint8) {
	if hex[0] == '#' {
		hex = hex[1:]
	}
	var r, g, b uint8
	fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	return r, g, b
}
