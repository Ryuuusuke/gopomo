package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	gopomo "github.com/ryuuusuke/gopomo/src"
)

func main() {
	fmt.Print("\033[H\033[2J")
	var timeInputFocus, timeInputRest time.Duration
	fmt.Printf("%s", gopomo.Cyan.Render("Insert Your Focus Time in Minute [25]: "))
	fmt.Scanln(&timeInputFocus)
	if timeInputFocus == 0 {
		timeInputFocus = 25
	}
	focusTime := time.Duration(timeInputFocus) * time.Minute

	fmt.Printf("%s", gopomo.Green.Render("Insert Your Rest Time in Minute [5]: "))
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
