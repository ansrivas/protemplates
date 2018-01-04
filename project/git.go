package project

import (
	"os/exec"
)

// InitIfGitExist initializes the given projectdir if git exists in the path
func InitIfGitExist(projectdir string) bool {
	_, err := exec.LookPath("git")
	if err != nil {
		return false
	}
	cmd := exec.Command("git", "init", projectdir)
	err = cmd.Run()
	if err != nil {
		return false
	}
	return true
}
