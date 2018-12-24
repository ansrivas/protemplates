package project

// Snippet taken from https://github.com/bclicn/color/blob/4c02eff8a28c1add99b3b5b238c165056c3bea8e/color.go
import (
	"strconv"
)

const (
	// common
	reset  = "\033[0m" // auto reset the rest of text to default color
	normal = 0
	bold   = 1 // increase this value if you want bolder text
	// special
	dim       = 2
	underline = 4
	blink     = 5
	reverse   = 7
	hidden    = 8
	// color
	black       = 30 // default = 39
	red         = 31
	green       = 32
	yellow      = 33
	blue        = 34
	purple      = 35 // purple = magenta
	cyan        = 36
	lightGray   = 37
	darkGray    = 90
	lightRed    = 91
	lightGreen  = 92
	lightYellow = 93
	lightBlue   = 94
	lightPurple = 95
	lightCyan   = 96
	white       = 97

	// SignSuccess represents that a task is done
	SignSuccess = "√"

	// SignFailure represents that a task is pending
	SignFailure = "x"
)

func render(colorCode int, fontSize int, content string) string {
	return "\033[" + strconv.Itoa(fontSize) + ";" + strconv.Itoa(colorCode) + "m" + content + reset
}

// RedText prints a given text in red color
func RedText(txt string) string {
	return render(red, normal, txt)
}

// GreenText prints a given text in red color
func GreenText(txt string) string {
	return render(green, normal, txt)
}
