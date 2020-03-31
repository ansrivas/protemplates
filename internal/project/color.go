//
// Copyright (c) 2020 Ankur Srivastava
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
