package src

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (model PomodoroModel) Init() tea.Cmd {
	return nil
}

func (model PomodoroModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.WindowSizeMsg:
		model.Width = msg.Width
		model.Height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return model, tea.Quit

		case "s":
			if !model.IsRunning {
				model.IsRunning = true
				return model, startTicking()
			}
		case "p":
			model.IsRunning = false

		case "r":
			model.RemainingTime = model.TotalDuration
			model.IsRunning = false

		}

	case tickMessage:
		if model.IsRunning {
			model.RemainingTime -= time.Second

			if model.RemainingTime <= 0 {
				model.RemainingTime = 0
				model.IsRunning = false
				return model, nil
			}
			return model, startTicking()
		}
	}
	return model, nil
}
