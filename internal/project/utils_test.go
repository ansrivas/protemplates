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

package project

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type utilsTestSuite struct {
	suite.Suite
	dirname  string
	testFile string
}

func (suite *utilsTestSuite) SetupTest() {
	suite.dirname = "existingDir"
	os.Mkdir(suite.dirname, os.ModePerm)
	log.Println("Setup")
}

func (suite *utilsTestSuite) TearDownTest() {
	os.RemoveAll(suite.dirname)
	log.Println("This should have been printed after each test to cleanup resoures.")
}

func (suite *utilsTestSuite) Test_CreateDir() {

	err := CreateDir("nonExistingDir")
	if suite.Nil(err, "Should successfully create a non-existing directory.") {
		os.RemoveAll("nonExistingDir")
	}
}

func (suite *utilsTestSuite) Test_MustCreateDir() {
	panicFunc := func() {
		MustCreateDir(suite.dirname)
	}
	suite.Panics(panicFunc, "MustCreateDir should panic on trying to recreate existing directory.")
}

func (suite *utilsTestSuite) Test_WriteToFile() {
	data := "Well Hello world"
	err := WriteToFile("testFile.txt", data)

	if suite.Nil(err, "File should be successfully written") {
		bytes, _ := ioutil.ReadFile("testFile.txt")
		if suite.Equal(data, string(bytes), "File contents should match after successful writing.") {
			// Apply the cleanup
			os.Remove("testFile.txt")
		}
	}
}

func TestIndexTestSuite(t *testing.T) {
	suite.Run(t, new(utilsTestSuite))
}

// ----------------------------------------------------------------------------------------------

func Test_SanitizeInput(t *testing.T) {
	assert := assert.New(t)
	testStruct := []struct {
		input  string
		output string
	}{
		{"pythoN", "python"},
		{"test-Project", "test-project"},
	}

	for _, s := range testStruct {
		assert.Equal(SanitizeInput(s.input), s.output, "Should properly sanitize input strings. ")
	}
}

// ----------------------------------------------------------------------------------------------

type TestLanguage struct{}

func (t TestLanguage) Create(appname string) error {
	fmt.Println("Call to create is successful with appname:", appname)
	return nil
}
func Test_Create(t *testing.T) {
	assert := assert.New(t)
	err := Create(TestLanguage{}, "myapp")
	assert.Nil(err, "Should successfully test creation of projects, using interface")
}

func Test_ParseTemplateString(t *testing.T) {
	assert := assert.New(t)
	tp := template.New("test")
	tplstr := `{{.input}}`
	expected := `testInput`
	actual := ParseTemplateString(tplstr, tp, map[string]string{"input": "testInput"})

	assert.Equal(expected, actual, "Template string parsing should work as expected.")
}
