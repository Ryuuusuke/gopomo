package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	gopomo "github.com/ryuuusuke/gopomo/src"
)

func main() {
	var timeInputFocus, timeInputRest time.Duration
	fmt.Print("Insert your Focus time in Minute [25]:")
	fmt.Scanln(&timeInputFocus)
	if timeInputFocus == 0 {
		timeInputFocus = 25
	}
	focusTime := time.Duration(timeInputFocus) * time.Minute

	fmt.Print("Insert your Rest time in Minute [5]:")
	fmt.Scanln(&timeInputRest)
	if timeInputRest == 0 {
		timeInputRest = 5
	}
	restTime := time.Duration(timeInputRest) * time.Minute

	initialModel := gopomo.PomodoroModel{
		FocusDuration: focusTime,
		RestDuration:  restTime,
		RemainingTime: focusTime,
		IsRunning:     false,
	}

	program := tea.NewProgram(initialModel)
	if _, err := program.Run(); err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

}
