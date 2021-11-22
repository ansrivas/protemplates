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

var makefileText = `.DEFAULT_GOAL := help

help:             ## Show available options with this Makefile
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -v grep | awk 'BEGIN { FS = ":.*?##" }; { printf "%-18s  %s\n", $$1,$$2 }'

.PHONY : test
test:             ## Run all the tests
	pytest

.PHONY : recreate_pyenv
recreate_pyenv:   ## Create the python environment. Recreates if the env exists already.
	conda env create --force -f dev_environment.yml

.PHONY : readme_to_rst
readme_to_rst:    ## Convert README.md to README.rst for the sake of pip documentation.
	m2r --overwrite README.md

.PHONY : upload_test_pypi
upload_test_pypi: readme_to_rst ## Build and upload distribution to testpypi server
	python setup.py bdist_wheel --dist-dir dist && \
	twine upload --skip-existing --repository testpypi dist/*

.PHONY : upload_pypi
upload_pypi:  readme_to_rst    ## Build and upload distribution to pypi server
	python setup.py bdist_wheel --dist-dir dist && \
	twine upload --skip-existing --repository testpypi dist/*

.PHONY : lint
lint:             ## Run both black and isort
	make black
	make isort

.PHONY : black
black:             ## Run black linter
	black {{.appWithUnderScore}}

.PHONY : isort
isort:             ## Run isort on the source directory
	isort --atomic --check-only --recursive {{.appWithUnderScore}}
`
