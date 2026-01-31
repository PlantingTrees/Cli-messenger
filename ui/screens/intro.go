package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/donderom/bubblon"
)

// gives a brief intro about what the app functionalities and feature request that links to the githu repo
type IntroModel struct {
	heading  string
	userName string
}

func NewIntroModel(name string) IntroModel {
	return IntroModel{
		heading:  "First of all, introduction...",
		userName: name,
	}
}

// tea

func (m IntroModel) Init() tea.Cmd {
	return nil
}

func (m IntroModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			return m, bubblon.Replace(NewChatModel())
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m IntroModel) View() string {
	return fmt.Sprintf("%s %s", m.heading, m.userName)
}
