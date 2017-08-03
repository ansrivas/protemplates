package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

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
	expected := "python"
	assert.Equal(SanitizeInput("pythoN"), expected, "Should properly sanitize input strings. ")
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
