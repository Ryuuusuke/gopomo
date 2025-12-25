package src

import (
	"time"
)

type tickMessage time.Time

type PomodoroModel struct {
	TotalDuration time.Duration
	RemainingTime time.Duration
	IsRunning     bool
	Width         int
	Height        int
}
