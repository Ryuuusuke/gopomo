package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (model pomodoroModel) Init() tea.Cmd {
	return nil
}

func (model pomodoroModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.WindowSizeMsg:
		model.width = msg.Width
		model.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return model, tea.Quit

		case "s":
			if !model.isRunning {
				model.isRunning = true
				return model, startTicking()
			}
		case "p":
			model.isRunning = false

		case "r":
			model.remainingTime = model.totalDuration
			model.isRunning = false

		}

	case tickMessage:
		if model.isRunning {
			model.remainingTime -= time.Second

			if model.remainingTime <= 0 {
				model.remainingTime = 0
				model.isRunning = false
				return model, nil
			}
			return model, startTicking()
		}
	}
	return model, nil
}
