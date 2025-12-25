package src

import (
	"time"
)

type tickMessage time.Time

type PomodoroModel struct {
	FocusDuration time.Duration
	RestDuration  time.Duration
	RemainingTime time.Duration
	IsRunning     bool
	IsResting     bool
	Width         int
	Height        int
}
