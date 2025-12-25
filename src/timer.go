package src

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func startTicking() tea.Cmd {
	return tea.Tick(time.Second, func(currentTime time.Time) tea.Msg {
		return tickMessage(currentTime)
	})
}
