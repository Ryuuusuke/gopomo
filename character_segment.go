package main

var colonSegments = [5]Segment{
	SegmentEmpty,
	SegmentCenter,
	SegmentEmpty,
	SegmentCenter,
	SegmentEmpty,
}

var numberSegments = [50]Segment{
	// 0
	SegmentFull, SegmentSides, SegmentSides, SegmentSides, SegmentFull,
	// 1
	SegmentRight, SegmentRight, SegmentRight, SegmentRight, SegmentRight,
	// 2
	SegmentFull, SegmentRight, SegmentFull, SegmentLeft, SegmentFull,
	// 3
	SegmentFull, SegmentRight, SegmentFull, SegmentRight, SegmentFull,
	// 4
	SegmentSides, SegmentSides, SegmentFull, SegmentRight, SegmentRight,
	// 5
	SegmentFull, SegmentLeft, SegmentFull, SegmentRight, SegmentFull,
	// 6
	SegmentFull, SegmentLeft, SegmentFull, SegmentSides, SegmentFull,
	// 7
	SegmentFull, SegmentRight, SegmentRight, SegmentRight, SegmentRight,
	// 8
	SegmentFull, SegmentSides, SegmentFull, SegmentSides, SegmentFull,
	// 9
	SegmentFull, SegmentSides, SegmentFull, SegmentRight, SegmentFull,
}
