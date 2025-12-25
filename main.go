package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	gopomo "github.com/ryuuusuke/gopomo/src"
)

func main() {
	initialDuration := 25 * time.Minute

	initialModel := gopomo.PomodoroModel{
		TotalDuration: initialDuration,
		RemainingTime: initialDuration,
		IsRunning:     false,
	}

	program := tea.NewProgram(initialModel)
	if _, err := program.Run(); err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

}
