package python

var setupyText = `from setuptools import setup, find_packages

setup(
    name='%s',
    version='0.1.0',
    include_package_data=True,
    install_requires=['future'],
    setup_requires=['pytest-runner', 'pytest'],
    tests_require=['pytest'],
    packages=find_packages(),
)`

var setupCfgText = `[aliases]
test=pytest

[tool:pytest]
addopts = --verbose -vv`

var conftestText = `# !/usr/bin/env python
# -*- coding: utf-8 -*-
"""Bunch of fixtures to be used across the tests."""

import pytest

@pytest.fixture(scope="function")
def test_fixture(request):
    """Create a test fixture."""

    myvar = 5

    def tear_down():
        # clean up here
        pass
    request.addfinalizer(tear_down)

    return myvar`

var testfileText = `# !/usr/bin/env python
# -*- coding: utf-8 -*-
"""Test modules."""


def test_list_files_older_than(test_fixture):
    """Run a test."""

    assert(5 == test_fixture)`

var initpyText = `__version__ = "0.1.0"`

var makefileText = `.DEFAULT_GOAL := help

help:          ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test
test:          ## Run all the tests
test:
	python setup.py test`
