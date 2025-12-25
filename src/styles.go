package src

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center)

	runningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42")).
			Bold(true).
			Align(lipgloss.Center)

	pausedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Align(lipgloss.Center)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Align(lipgloss.Center)

	focusColor = lipgloss.Color("205")
	restColor  = lipgloss.Color("#58fc97")
)
