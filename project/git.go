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
