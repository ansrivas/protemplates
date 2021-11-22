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

var pyprojectText = `[aliases]
test=pytest

[tool:pytest]
addopts = --verbose -vv --cov-report term-missing --cov {{.appWithUnderScore}}

[tool.black]
line-length = 120
target-version = ['py37', 'py38', 'py39']
include = '\.pyi?$'
extend-exclude = '''
'''

[isort]
line_length = 120
skip = __init__.py, setup.py
indent = '    '
multi_line_output = 0
length_sort = 0

[tool.poetry]
name = "{{.appWithUnderScore}}"
version = "5.10.1"
description = "Describe your project here."
authors = ["{{.author}} <{{.authoremail}}>"]
license = "MIT"
readme = "README.md"
repository = "https://github.com/{{.scmusername}}/{{.appname}}"
classifiers = [
    "Development Status :: 6 - Mature",
    "Intended Audience :: Developers",
    "Natural Language :: English",
    "Environment :: Console",
    "License :: OSI Approved :: MIT License",
    "Programming Language :: Python",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.6",
    "Programming Language :: Python :: 3.7",
    "Programming Language :: Python :: 3.8",
    "Programming Language :: Python :: 3.9",
    "Programming Language :: Python :: 3.10",
    "Programming Language :: Python :: 3 :: Only",
    "Programming Language :: Python :: Implementation :: CPython",
    "Programming Language :: Python :: Implementation :: PyPy",
    "Topic :: Software Development :: Libraries",
    "Topic :: Utilities",
]
urls = { Changelog = "https://github.com/{{.author}}/{{.appname}}/blob/main/CHANGELOG.md" }
packages = [
    { include = "{{.appWithUnderScore}}" },
]
include = [
    { path = "tests", format = "sdist" },
]


[tool.poetry.dependencies]
python = ">=3.6.2,<4.0"

[tool.poetry.dev-dependencies]
black = {version = "^21.10b0", allow-prereleases = true}
coverage = {version = "^6.0b1", allow-prereleases = true}
mypy = "^0.902"
pytest = "^6.0"
pytest-cov = "^2.7"
pytest-mock = "^1.10"


[tool.poetry.scripts]
{{.appWithUnderScore}} = "{{.appWithUnderScore}}.main:run"

[build-system]
requires = ["poetry-core>=1.0.0"]
build-backend = "poetry.core.masonry.api"

`
