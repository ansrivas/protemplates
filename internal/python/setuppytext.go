// MIT License
//
// Copyright (c) 2019 Ankur Srivastava
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

var setupyText = `import re
from os import path
from codecs import open  # To use a consistent encoding

from setuptools import setup, find_packages

here = path.abspath(path.dirname(__file__))

# Get the long description from the relevant file
with open(path.join(here, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()


# Get version without importing, which avoids dependency issues
def get_version():
    with open('{{.appWithUnderScore}}/__init__.py') as version_file:
        return re.search(r"""__version__\s+=\s+(['"])(?P<version>.+?)\1""",
                         version_file.read()).group('version')


install_requires = ['future']


TEST_REQUIRES = [
    "pytest==5.2.1",
    "pytest-sugar==0.9.2",
    "pytest-asyncio==0.10.0",
    "pytest-cov==2.8.1",
]


EXTRAS_REQUIRE = {
    "dev": ["python-language-server[all]", "black==19.3b0", "isort"],
    "test": TEST_REQUIRES,
}


setup(
    name='{{.appWithHyphen}}',
    description="Some description about your project",
    long_description=long_description,
    long_description_content_type='text/markdown',
    version=get_version(),
    include_package_data=True,
    install_requires=install_requires,
    setup_requires=['pytest-runner'],
    extras_require=EXTRAS_REQUIRE,
    tests_require=TEST_REQUIRES,
    packages=find_packages(),
    zip_safe=False,
    author="{{.author}}",
    author_email = "{{.authoremail}}",
    download_url="https://{{.scm}}/{{.scmusername}}/{{.appWithHyphen}}/{}.tar.gz".format(get_version()),
    classifiers=[
        "Programming Language :: Python :: 2",
        "Programming Language :: Python :: 2.7",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.4",
        "Programming Language :: Python :: 3.5",
        "Programming Language :: Python :: 3.6", ]
)
`
