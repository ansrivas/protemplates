package golang

import (
	"fmt"
	"path"

	"github.com/ansrivas/protemplates/internal"
)

// Golang ...
type Golang struct {
	Scm      string
	Username string
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
		internal.MustCreateDir(temp)
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
		err := internal.WriteToFile(path, content)
		if err != nil {
			return fmt.Errorf("Unable to write file: %s", path)
		}
	}

	return nil
}
