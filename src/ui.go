package src

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func formatRemainingTime(duration time.Duration) string {
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func renderTitle(model PomodoroModel) string {
	focusTitle := titleStyle.Foreground(focusColor).Render("Go Pomo!")
	if model.IsResting == true {
		resttitle := titleStyle.Foreground(restColor).Render("Go Rest!")
		return resttitle
	}
	return focusTitle
}

func (model PomodoroModel) View() string {
	// format time
	timeText := formatRemainingTime(model.RemainingTime)
	bigClock := renderClockTime(timeText, ColorBlue)
	title := renderTitle(model)

	// status
	statusText := "⏸ PAUSED"
	statusView := pausedStyle.Render(statusText)
	if model.IsRunning {
		statusText = "▶ RUNNING"
		statusView = runningStyle.Render(statusText)
	}

	content := strings.Join([]string{
		title,
		"",
		bigClock,
		"",
		statusView,
		"",
		helpStyle.Render("[s] start • [p] pause • [r] reset • [q] quit"),
	}, "\n")

	return lipgloss.Place(
		model.Width,
		model.Height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
