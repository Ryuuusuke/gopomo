package main

import (
	"time"
)

type tickMessage time.Time

type pomodoroModel struct {
	totalDuration time.Duration
	remainingTime time.Duration
	isRunning     bool
	width         int
	height        int
}
