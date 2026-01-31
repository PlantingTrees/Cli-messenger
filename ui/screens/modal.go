package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Modal struct {
	showModal bool
	yesFocus  bool // true = Yes, false = No
}

func NewModal() Modal {
	return Modal{
		showModal: false,
		yesFocus:  false, // Default to No
	}
}

func (m Modal) Value() bool {
	return m.showModal
}

func (m Modal) Update(msg tea.Msg) (Modal, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.showModal {
			switch msg.String() {
			case "left", "right", "tab":
				m.yesFocus = !m.yesFocus
				return m, nil
			case "enter":
				if m.yesFocus {
					return m, tea.Quit
				}
				m.showModal = false
				return m, nil
			case "y":
				return m, tea.Quit
			case "n", "esc":
				m.showModal = false
				return m, nil
			}
			return m, nil
		}

		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.showModal = true
			return m, nil
		}
	}
	return m, nil
}

func (m Modal) View() string {
	if !m.showModal {
		return ""
	}

	// Styles for buttons
	activeStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#FF0055")).
		Padding(0, 1).
		Bold(true)

	inactiveStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Padding(0, 1)

	var yes, no string
	if m.yesFocus {
		yes = activeStyle.Render("I GOTTA GO!")
		no = inactiveStyle.Render("ok, i stay")
	} else {
		yes = inactiveStyle.Render("I GOTTA GO!")
		no = activeStyle.Render("ok, i stay")
	}

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, yes, "  ", no)

	ui := lipgloss.JoinVertical(
		lipgloss.Center,
		"Don't go! 😢",
		"\n",
		buttons,
	)

	return lipgloss.NewStyle().
		Width(40).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0055")).
		Padding(1, 0).
		Align(lipgloss.Center).
		Render(ui)
}
