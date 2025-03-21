package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

//view of the Model and repr of the string

func (m Model) View() string {
	s := "what do you want to do with wrap today ?\n\n"

	if m.quitting {
		return ""
	}

	//iterate on choices
	for i, choice := range m.choices {
		//is curosr pointing at this choice ?
		current := " "
		if m.current == i {
			current = ">"
		}

		//is selected ?
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
			var s strings.Builder
			s.WriteString("\n  ")
			if m.err != nil {
				s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
			} else if m.selectedFile == "" {
				s.WriteString("Pick a file:")
			} else {
				s.WriteString("Selected file: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
			}
			s.WriteString("\n\n" + m.filepicker.View() + "\n")
			return s.String()
		}

		//render row
		s += fmt.Sprintf("%s [%s] %s\n", current, checked, choice)
	}

	//footer
	s += "\nPress q to quit.\n"

	//render the whole string for UI
	return s
}

func InitialModel() Model {
	return Model{
		choices: []string{"Generate Overview", "Make a graphical view using Graphviz", "Suggest new ideas based on existing ideas"},
		//map which indicates which choices are selected
		selected: make(map[int]struct{}),
	}
}

// init can return a Cmd that could perform some initial I/O.
func (m Model) Init() tea.Cmd {
	fmt.Printf("warp v.0.0.1\nSummarizing tool for large codebases\n")
	return m.filepicker.Init()
}
