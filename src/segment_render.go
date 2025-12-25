package src

const segmentWidth = 6

func (s Segment) Render(color Color) string {
	bg := color.Background()
	reset := ColorReset

	switch s {
	case SegmentFull:
		return bg + "      " + reset
	case SegmentLeft:
		return bg + "  " + reset + "    "
	case SegmentCenter:
		return "  " + bg + "  " + reset + "  "
	case SegmentRight:
		return "    " + bg + "  " + reset
	case SegmentSides:
		return bg + "  " + reset + "  " + bg + "  " + reset
	case SegmentEmpty:
		return "      "
	default:
		return "      "
	}
}
