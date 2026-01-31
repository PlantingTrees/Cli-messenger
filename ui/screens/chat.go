package screens

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/donderom/bubblon"
	"github.com/plantingtrees/cli-messenger/ui/components"
)

type ChatModel struct {
	chatBox textinput.Model
}

func NewChatModel() ChatModel {
	ti := components.NewInput()

	return ChatModel{
		chatBox: ti,
	}
}

func (m ChatModel) Init() tea.Cmd {
	return nil
}

func (m ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyShiftTab:
			return m, bubblon.Open(NewIntroModel("")) // this for reentry, im gonna fix this!
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m ChatModel) View() string {
	return "This is chat lobby "
}
