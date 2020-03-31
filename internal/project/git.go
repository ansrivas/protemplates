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
	"os"
	"os/exec"
)

// InitIfGitExist initializes the given projectdir if git exists in the path
func InitIfGitExist(projectdir string) bool {
	_, err := exec.LookPath("git")
	if err != nil {
		return false
	}
	// Check if directory already exists. Panic if it is already there
	if _, err = os.Stat(projectdir); err == nil {
		panic(fmt.Sprintf("Directory: %s already exists!!", projectdir))
	}

	cmd := exec.Command("git", "init", projectdir)
	err = cmd.Run()
	if err != nil {
		return false
	}
	return true
}

// InitialCommit creates the first commit with commmit message "first-commit"
func InitialCommit(projectdir string) bool {
	_, err := exec.LookPath("git")
	if err != nil {
		return false
	}
	// Check if directory doesn't exists. Panic.
	if _, err = os.Stat(projectdir); os.IsNotExist(err) {
		panic(fmt.Sprintf("Directory: %s does not exist!!", projectdir))
	}

	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s && git add . && git commit -m 'initial-commit'", projectdir))
	err = cmd.Run()
	if err != nil {
		return false
	}
	return true
}
