package main

import (
	"fmt"
	"os"
	"time"

	"github.com/plantingtrees/cli-messenger/ui"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define a message for the tick
type TickMsg time.Time

// --- Model ---
type model struct {
	textInput textinput.Model
	width     int
	height    int
	frame     int
}

func initialModel() model {
	return model{
		textInput: ui.NewInput(),
		frame:     5,
	}
}

// --- Init ---
func (m model) Init() tea.Cmd {
	// Return a batch: Blink the cursor AND start the animation timer
	return tea.Batch(
		textinput.Blink,
		tick(),
	)
}

// --- Update ---
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// Capture the Tick
	case TickMsg:
		m.frame++        // Increment the frame
		return m, tick() // Trigger the next tick immediately

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			val := m.textInput.Value()
			return m, tea.Sequence(
				tea.Println("OYEZ Protocol Initiated: "+val),
				tea.Quit,
			)
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// --- View ---
func (m model) View() string {
	// PASS THE FRAME TO THE LOGO
	logo := ui.RenderLogo(m.frame)

	inputBox := ui.RenderInput(m.textInput)
	helpText := ui.RenderHelp("ESC TO ABORT • ENTER TO TRANSMIT")

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		logo,
		inputBox,
		helpText,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

// --- Helper for the Tick ---
func tick() tea.Cmd {
	// Update every 17 milliseconds
	return tea.Tick(time.Millisecond*17, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
