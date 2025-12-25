package src

type Segment int

const (
	SegmentFull Segment = iota
	SegmentLeft
	SegmentCenter
	SegmentRight
	SegmentSides
	SegmentEmpty
)
