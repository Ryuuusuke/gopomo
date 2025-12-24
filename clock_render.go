package main

import "strings"

func renderClockTime(timeText string, color Color) string {
	chars := parseTimeToCharacters(timeText)
	lines := make([]string, 5)

	for row := 0; row < 5; row++ {
		var line strings.Builder

		for _, ch := range chars {
			line.WriteString(ch.RenderRow(row, color))
			line.WriteString(" ")
		}

		lines[row] = line.String()
	}

	return strings.Join(lines, "\n")
}
