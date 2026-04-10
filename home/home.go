package home

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	width, height int
	title         string
	menuItems     []string
}

func NewModel(width, height int, title string, items []string) model {
	return model{
		width:     width,
		height:    height,
		title:     title,
		menuItems: items,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() tea.View {
	content := fmt.Sprintf("%s\n\n", m.title)

	for _, item := range m.menuItems {
		content += fmt.Sprintf("%s", item)
	}

	return tea.NewView(content)
}
