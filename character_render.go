package main

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
