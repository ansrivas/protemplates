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

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ansrivas/protemplates/internal/golang"
	"github.com/ansrivas/protemplates/internal/licenses"
	"github.com/ansrivas/protemplates/internal/project"
	"github.com/ansrivas/protemplates/internal/python"
)

const (
	// readAuthorName is a string constant to be used as a prompt.
	readAuthorName = "Please enter author name: for eg. Lovan Vivan"

	// readAuthorEmail is a string constant to be used as a prompt.
	readAuthorEmail = "Please enter author email:"

	// readSCMInfo is a string constant to be used as a prompt.
	readSCMInfo = "Please enter your scm for eg. github.com, bitbucket.com"

	// readSCMUserName is a string constant to be used as a prompt.
	readSCMUserName = "Please enter a username corresponding to your scm eg. github.com/ansrivas, then ansrivas"
)

// mustCreateProject tries to create a project and panics if something fails.
func mustCreateProject(impl project.Project, projectName string) {
	err := project.Create(impl, projectName)
	if err != nil {
		panic(fmt.Sprintf("Unable to create the project: %s", projectName))
	}
	fmt.Printf("Successfully created the project %s in current directory\n", projectName)
}

// userAgrees checks if user agrees to a given query
func userAgrees(msg string) bool {
	var input string

	fmt.Printf("%s [y/N]: ", msg)
	for {
		fmt.Scanf("%s", &input)
		switch input {
		case "y":
			return true
		case "N":
			return false
		default:
			fmt.Println("Valid responses are only [y/N]")
			continue
		}
	}
}

func readInput(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg)
	var input string
	for {
		scanner.Scan()
		input = scanner.Text()
		if input == "" && !userAgrees("Continuing with empty input") {
			continue
		}
		break
	}
	return input
}

func createProject(language string) {
	lang := project.SanitizeInput(language)

	var projectName string
	fmt.Println("Please enter a desired project name:")
	for {
		fmt.Scanf("%s", &projectName)
		if projectName != "" {
			break
		}
		fmt.Println("Project name can not be empty, please enter a valid project name:")
	}

	fmt.Println(`Please enter a license type.
[1] MIT
[2] Apache2`)

	var license string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		license = ""
		scanner.Scan()
		switch scanner.Text() {
		case "1":
			license = licenses.MIT
		case "2":
			license = licenses.Apache2
		default:
			fmt.Println("Licenses can be either MIT or Apache2 currently")
		}
		if license != "" {
			break
		}
	}

	author := readInput(readAuthorName)
	authoremail := readInput(readAuthorEmail)
	scm := readInput(readSCMInfo)
	scmusername := readInput(readSCMUserName)

	var impl project.Project
	switch lang {
	case "python":
		impl = python.New(projectName, license, author, authoremail, scm, scmusername)
	case "go", "golang":
		impl = golang.New(projectName, license, author, authoremail, scm, scmusername)
	default:
		fmt.Printf("\033[31mException: Language %s is currently not supported.\033[39m\n\n", language)
		os.Exit(1)
	}
	mustCreateProject(impl, projectName)
}
