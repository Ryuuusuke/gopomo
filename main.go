package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initialDuration := 25 * time.Minute

	initialModel := pomodoroModel{
		totalDuration: initialDuration,
		remainingTime: initialDuration,
		isRunning:     false,
	}

	program := tea.NewProgram(initialModel)
	if _, err := program.Run(); err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

}
