package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	column  = lipgloss.NewStyle().Padding(1, 2)
	focused = lipgloss.NewStyle().Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))
)

func main() {
	if err := tea.NewProgram(model{}).Start(); err != nil {
		log.Fatal(err)
	}
}

type model struct{}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		// TODO(jay): How in the hell do you resize a window?
	}
	return m, nil
}

func (m model) View() string {
	return focused.Render("Look at me")
}
