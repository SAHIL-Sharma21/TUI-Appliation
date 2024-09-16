package main

import (
	"fmt"

	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

// models
type models struct {
	choices  []string         //list of items
	cursor   int              //selecting the items
	selected map[int]struct{} //seleted items
}

func initialModel() models {
	return models{
		choices:  []string{"Golang Bootcamp", "MERN Bootcamp", "System Design", "Devops Bootcamp", "DSA + interview Prep"},
		selected: make(map[int]struct{}),
	}
}

func (m models) Init() tea.Cmd {
	return nil
}

func (m models) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	//key pressed
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m models) View() string {
	s := "Which course should I take?\n\n"

	color.Set(color.FgMagenta)
	s += fmt.Sprintf("%s\n", m.choices[m.cursor])

	for i, choice := range m.choices {
		cursor := " " //no cursor

		if m.cursor == i {
			cursor = ">"
		}

		//choice is seleted
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "X"
		}

		//render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\n(Press q to quit)"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Alas, there's been an error:", err)
		os.Exit(1)
	}
}
