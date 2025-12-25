package src

import (
	"time"
)

type tickMessage time.Time

type PomodoroModel struct {
	totalDuration time.Duration
	remainingTime time.Duration
	isRunning     bool
	width         int
	height        int
}
