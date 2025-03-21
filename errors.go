package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

//utiliy helpful things and error handling

type ClearErrorMsg struct{}

func ClearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return ClearErrorMsg{}
	})
}
