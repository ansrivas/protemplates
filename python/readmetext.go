package python

const (
	ticksthree  = "```"
	newlineone  = "\n"
	newlinestwo = "\n\n"
	tabone      = "\t"
	tabstwo     = "\t\t"
	tabsthree   = "\t\t\t"
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
		commandMarkdown(`$ pip install -e .`) + newlinestwo +
		`## Test` + newlineone +
		`To run the tests:` + newlineone +
		commandMarkdown(`make test`) + newlineone +
		`## Usage` + newlineone +
		`## Examples` + newlineone +
		commandMarkdown(`$ python examples/simple.py`) + newlineone +
		"## License" + newlineone +
		"{{.license}}"
}
