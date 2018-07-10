package python

var activateEnvText = `source activate {{.appWithHyphen}}-env`

var setupCfgText = `[aliases]
test=pytest

[tool:pytest]
addopts = --verbose -vv --cov-report term-missing --cov {{.appWithUnderScore}}

[pep8]
max-line-length=119

[flake8]
ignore = N801,N802,N803,W503,E12
max-line-length=119
`

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
- python=3.6.5
- pip:
  - flake8
  - autopep8
  - yapf
  - cython
  - future
  - twine
  - m2r
`
