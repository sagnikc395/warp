package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type DataModel struct {
	choices  []string
	current  int
	selected map[int]struct{}
}

func initialModel() DataModel {
	return DataModel{
		choices: []string{"Generate Overview", "Make a graphical view using Graphviz", "Suggest new ideas based on existing ideas"},
		//map which indicates which choices are selected
		selected: make(map[int]struct{}),
	}
}

// init can return a Cmd that could perform some initial I/O.
func (dm DataModel) Init() tea.Cmd {
	fmt.Printf("warp v.0.0.1\nSummarizing tool for large codebases\n")

	return nil
}

//well figure out which type of Msg we received with a type switch
// rn only dealing with tea.KeyMsg messages, whcih are automatically sent to the update
// function when he keys are pressed

func (m DataModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	//key press ?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.current > 0 {
				m.current--
			}

		case "down", "j":
			if m.current < len(m.choices)-1 {
				m.current++
			}

		//enter key and the spacebar -> the selected state for the item under it
		case "enter", " ":
			_, ok := m.selected[m.current]
			if ok {
				delete(m.selected, m.current)
			} else {
				m.selected[m.current] = struct{}{}
			}
		}
	}

	//return the updated mode to bl runtime
	return m, nil
}

// view renders our UI.we look at the UI in the current state and use it to return a string.
// the string is out UI
func (m DataModel) View() string {
	s := "what do you want to do with wrap today ?\n\n"

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
		}

		//render row
		s += fmt.Sprintf("%s [%s] %s\n", current, checked, choice)
	}

	//footer
	s += "\nPress q to quit.\n"

	//render the whole string for UI
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}
}
