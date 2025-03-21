package main

import (
	"errors"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

//well figure out which type of Msg we received with a type switch
// rn only dealing with tea.KeyMsg messages, whcih are automatically sent to the update
// function when he keys are pressed

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	//key press ?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
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
	case ClearErrorMsg:
		m.err = nil
	}

	//filepicker
	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	// did the user select a file ?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		m.selectedFile = path
	}

	//did the user select a disable file ?
	//only necessary to display a error to the user
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, ClearErrorAfter(2*time.Second))
	}

	//return the updated mode to bl runtime
	return m, cmd
}
