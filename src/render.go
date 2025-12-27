package src

import "strings"

func (c Character) RenderRow(row int, color Color) string {
	var segment Segment

	switch c.kind {
	case CharNumber:
		segment = numberSegments[c.value*5+row]
	case CharColon:
		segment = colonSegments[row]
	default:
		segment = SegmentEmpty
	}

	return segment.Render(color)
}

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
