package golang

import (
	"fmt"
	"path"

	"github.com/ansrivas/protemplates/project"
)

// Golang struct is responsible for handling golang projects
type Golang struct {
	Scm      string
	Username string
}

// New creates a new implementation for golang which is later used to create a project.
func New(projectName string) project.Project {
	var scm, username string

	fmt.Println("Please enter a desired scm eg. github.com or bitbucket.com")
	fmt.Scanf("%s", &scm)

	fmt.Println("Please enter a username corresponding to your scm eg. github.com/ansrivas, then ansrivas")
	fmt.Scanf("%s", &username)

	if username == "" || scm == "" {
		fmt.Println("Username or scm can not be empty")
		return nil
	}

	return Golang{
		Scm:      scm,
		Username: username,
	}
}

// Create creates a template folder structure for a golang project.
func (g Golang) Create(appname string) error {

	basedir := "src"
	srcdir := path.Join(appname, basedir, g.Scm, g.Username, appname)

	//order matters.
	dirs := []string{appname, basedir, g.Scm, g.Username, appname}
	temp := ""
	for _, dir := range dirs {
		temp = path.Join(temp, dir)
		project.MustCreateDir(temp)
	}

	pathToContent := make(map[string]string)

	makefilePath := path.Join(srcdir, "Makefile")
	readmePath := path.Join(srcdir, "README.md")
	testShellPath := path.Join(srcdir, "test.sh")
	gitignorePath := path.Join(srcdir, ".gitignore")
	mainPath := path.Join(srcdir, "main.go")

	pathToContent[makefilePath] = fmt.Sprintf(makefileText, appname, g.Username, appname)
	pathToContent[readmePath] = fmt.Sprintf(readmeText, appname)
	pathToContent[testShellPath] = testShellText
	pathToContent[gitignorePath] = gitignoreText
	pathToContent[mainPath] = mainText

	for path, content := range pathToContent {
		err := project.WriteToFile(path, content)
		if err != nil {
			return fmt.Errorf("Unable to write file: %s", path)
		}
	}

	return nil
}
