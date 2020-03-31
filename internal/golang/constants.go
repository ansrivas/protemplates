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

package golang

var readmeText = `{{.appname}}:
	---

	This project can be used to ...

	Install:
	---
	Clone the project and run ...

	Test:
	---
	To run the tests: ...


	Usage:
	---

	....

	Example:
	---

	...
	`

var mainText = `package main

import (
	"fmt"
)

var (
	// BuildTime gets populated during the build proces
	BuildTime = ""

	//Version gets populated during the build process
	Version = ""
)

func helloWorld() string{
	return "Hello World"
}

func main() {
	fmt.Printf("Current version is: %s and buildtime is: %s\n", Version, BuildTime)
}
`

var mainTestText = `package main

import(
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_helloWorld(t *testing.T) {
	assert := assert.New(t)
	expected:= "Hello World"
	assert.Equal( expected, helloWorld(), "Failed test HelloWorld")
}
`
