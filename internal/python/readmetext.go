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

package python

const (
	ticksthree  = "```"
	newlineone  = "\n"
	newlinestwo = "\n\n"
)

// commandMarkdown returns a markdown formatted version of a given input command with ticks, etc.
func commandMarkdown(cmd string) string {
	return ticksthree + newlineone +
		cmd + newlineone +
		ticksthree
}

var readmeText = func() string {
	return `# {{.appWithHyphen}}
[![Build Status](https://travis-ci.org/username/%s.svg?branch=master)](https://travis-ci.org/username/%s)

Short Description.

## Current Stable Version` + newlineone +
		commandMarkdown("0.1.0") + newlineone +
		`## Installation` + newlineone +
		`### pip` + newlineone +
		commandMarkdown(`pip install {{.appWithHyphen}}`) + newlineone +
		`### Development Installation` + newlineone +
		`* Clone the project.` + newlineone +
		`* Install in Anaconda3 environment` + newlineone +
		`* This command creates a python environment and then activates it.` + newlineone +
		commandMarkdown(`$ make recreate_pyenv && chmod +x activate-env.sh && . activate-env.sh`) + newlineone +
		`* Now install the application in editable mode and you are ready to start development` + newlineone +
		commandMarkdown(`$ pip install -e .[dev]`) + newlinestwo +
		`## Test` + newlineone +
		`To run the tests:` + newlineone +
		commandMarkdown(`make test`) + newlineone +
		`## Usage` + newlineone +
		`## Examples` + newlineone +
		commandMarkdown(`$ python examples/simple.py`) + newlineone +
		"## License" + newlineone +
		"{{.license}}"
}
