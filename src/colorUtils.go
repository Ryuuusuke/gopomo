package src

type Color string

const (
	ColorReset string = "\033[0m"
	ColorRed   Color  = "\033[41m"
	ColorGreen Color  = "\033[42m"
	ColorBlue  Color  = "\033[44m"
)

func (c Color) Background() string {
	return string(c)
}
