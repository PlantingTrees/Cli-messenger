package ui

import "github.com/charmbracelet/lipgloss"

var helpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240")).
	MarginTop(1)

func RenderHelp(text string) string {
	return helpStyle.Render(text)
}
