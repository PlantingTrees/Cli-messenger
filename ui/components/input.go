package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// We define the style here so it's reusable
var containerStyle = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("#FF0055")).
	Padding(0, 1).
	Width(40).
	MarginTop(2) // Keeps the spacing between logo and box

// NewInput creates a correctly configured text input
func NewInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter Username"
	ti.Focus()
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#813452"))
	ti.CharLimit = 20
	ti.Width = 30
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0055"))
	return ti
}

// RenderInput wraps the textinput.View() in our custom border
func RenderInput(ti textinput.Model) string {
	return containerStyle.Render(ti.View())
}
