package python

const (
	ticksthree    = "```"
	newlineone    = "\n"
	newlinestwo   = "\n\n"
	newlinesthree = "\n\n\n"
	tabone        = "\t"
	tabstwo       = "\t\t"
	tabsthree     = "\t\t\t"
)

var readmeText = `# {{.appWithHyphen}}
[![Build Status](https://travis-ci.org/username/%s.svg?branch=master)](https://travis-ci.org/username/%s)

Short Description.

#### Current Stable Version` + newlineone +
	ticksthree + newlineone +
	"0.1.0" + newlineone +
	ticksthree + newlinesthree +

	`
# Installation

### pip` + newlineone +
	ticksthree + newlineone +
	`pip install %s` + newlineone +
	ticksthree + newlinesthree +

	`### Development Installation

* Clone the project.
* Install in Anaconda3 environment` + newlineone +
	ticksthree + newlineone +
	`$ conda env create --force -f dev_environment.yml
$ source activate {{.appname}}
$ pip install -e .` + newlineone +
	ticksthree + newlinesthree +
	`# Test

To run the tests:` + newlineone +
	ticksthree + newlineone +
	`make test` + newlineone +
	ticksthree + newlinesthree +

	`# Usage


# Examples` + newlineone +
	ticksthree + newlineone +
	`$ python examples/simple.py` + newlineone +
	ticksthree + newlinesthree +
	"# License" + newlineone +
	"{{.license}}"
