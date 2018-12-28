package python

var setupCfgText = `[aliases]
test=pytest

[tool:pytest]
addopts = --verbose -vv --cov-report term-missing --cov {{.appWithUnderScore}}

[pep8]
max-line-length=120

[isort]
line_length = 120
skip = __init__.py, setup.py
indent = '    '
multi_line_output = 0
length_sort = 0

[flake8]
ignore = D203
exclude =
    # No need to traverse our git directory
    .git,
    # There's no value in checking cache directories
    __pycache__,
    # The conf file is mostly autogenerated, ignore it
    docs/source/conf.py,
    # The old directory contains Flake8 2.0
    old,
    # This contains our built documentation
    build,
    # This contains builds of flake8 that we don't want to check
    dist
    # Ignore F401 flake8
    __init__.py
max-complexity = 10
max-line-length=120`
