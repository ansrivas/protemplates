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

var activateEnvText = `source activate {{.appWithHyphen}}-env`

var conftestText = `# !/usr/bin/env python
# -*- coding: utf-8 -*-
"""Bunch of fixtures to be used across the tests."""

import pytest


@pytest.fixture(scope="function")
def hello_world(request):
    """Create a test fixture."""
    hw = "Hello World!"

    def tear_down():
        # clean up here
        pass

    request.addfinalizer(tear_down)
    return hw
`

var testfileText = `# !/usr/bin/env python
# -*- coding: utf-8 -*-
"""Test modules."""


def test_init(hello_world):
    """Run a test."""
    import {{.appWithUnderScore}}

    # Test __init__
    assert hasattr({{.appWithUnderScore}}, '__version__')

    # Test pytest fixtures
    assert(hello_world == "Hello World!")
`

var initpyText = `__version__ = "0.1.0"
`

var requirementsText = `
`

var manifestText = `include README.md
`

var devEnvYamlText = `name: {{.appWithHyphen}}-env
channels:
  - defaults
dependencies:
- python=3.6.7
- pip:
  - flake8
  - autopep8
  - yapf
  - cython
  - future
  - twine
  - m2r
  - isort
`
