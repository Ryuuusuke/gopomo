package src

import (
	"os/exec"
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
			model.RemainingTime = model.FocusDuration
			model.IsRunning = false
			model.IsResting = false

		}

	case tickMessage:
		if model.IsRunning {
			model.RemainingTime -= time.Second

			if model.RemainingTime <= 0 {
				model.IsRunning = false

				if !model.IsResting {
					model.IsResting = true
					model.RemainingTime = model.RestDuration
					renderTitle(model) // render title to Go Rest!
					exec.Command("notify-send", "Gopomo", "Hey, It's time to Go Rest! come and start your Rest Timer!").Run()
				} else {
					model.IsResting = false
					model.RemainingTime = model.FocusDuration
					renderTitle(model) // back to Go Pomo!
					exec.Command("notify-send", "Gopomo", "Hi, Your Rest time is just ended, let's get Productive! come and start your focus timer!").Run()
				}

			}
			return model, startTicking()
		}
	}
	return model, nil
}
