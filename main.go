package main

import (
	"log"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/donderom/bubblon"
	"github.com/plantingtrees/cli-messenger/ui/components"
	"github.com/plantingtrees/cli-messenger/ui/screens"
)

// to animate, logo
// uses a custom message that gets sent to update()
type tickMsg time.Time

// tea run a go routine on the anon function, pauses fucntion for 10ms, and send the time to Update
func tick() tea.Cmd {
	return tea.Tick(30*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type model struct {
	frame         int
	logo          string
	userName      textinput.Model
	help          string
	showExitModal screens.Modal

	// for view rendering
	width  int
	height int
}

func NewModel() model {
	ti := components.NewInput()
	return model{
		frame:         0,
		logo:          "",
		userName:      ti,
		help:          "",
		showExitModal: screens.NewModal(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("OYEH"),
		textinput.Blink,
		tick(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// msg for logo animation
	case tickMsg:
		m.frame++
		return m, tick()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.userName.Value() != "" {
				return m, bubblon.Replace(screens.NewIntroModel(m.userName.Value()))
			}
		}
	}

	// exit modal handles its own logic
	var cmd tea.Cmd
	m.showExitModal, cmd = m.showExitModal.Update(msg)

	// prevents shadow typing in the textbox
	var inputCmd tea.Cmd
	if !m.showExitModal.Value() {
		m.userName, inputCmd = m.userName.Update(msg)
	}
	return m, tea.Batch(cmd, inputCmd)
}

func (m model) View() string {
	logoView := components.RenderLogo(m.frame)
	textBoxView := components.RenderInput(m.userName)
	helpView := components.RenderHelp("Esc to quit • Enter to continue")

	layoutElements := lipgloss.JoinVertical(
		lipgloss.Center,
		logoView,
		textBoxView,
		helpView,
	)

	// Base background layout
	view := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		layoutElements,
	)

	// IF MODAL IS OPEN: Overlay it
	if m.showExitModal.Value() {
		modalView := m.showExitModal.View()

		// This places the modal precisely in the center of the terminal
		return lipgloss.Place(m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			modalView,
			// This makes the background layout stay behind it
			lipgloss.WithWhitespaceChars(""),
			lipgloss.WithWhitespaceForeground(lipgloss.Color("0")),
		)
	}

	return view
}

func main() {
	controller, err := bubblon.New(NewModel())
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(controller, tea.WithAltScreen())
	m, _ := p.Run()

	if m, ok := m.(bubblon.Controller); ok && m.Err != nil {
		log.Fatal(m.Err)
	}
}
