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

func (model PomodoroModel) View() string {
	// format time
	timeText := formatRemainingTime(model.remainingTime)
	bigClock := renderClockTime(timeText, ColorBlue)

	// status
	statusText := "⏸ PAUSED"
	statusView := pausedStyle.Render(statusText)
	if model.isRunning {
		statusText = "▶ RUNNING"
		statusView = runningStyle.Render(statusText)
	}

	content := strings.Join([]string{
		titleStyle.Render("Go Pomo!"),
		"",
		bigClock,
		"",
		statusView,
		"",
		helpStyle.Render("[s] start • [p] pause • [r] reset • [q] quit"),
	}, "\n")

	return lipgloss.Place(
		model.width,
		model.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
