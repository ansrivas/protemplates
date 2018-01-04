package golang

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"text/template"
	"time"

	"github.com/ansrivas/protemplates/licenses"
	"github.com/ansrivas/protemplates/project"
)

// Golang struct is responsible for handling golang projects
type Golang struct {
	project.Info
}

var t = template.New("golang")

// New creates a new implementation for golang which is later used to create a project.
func New(projectName, license, author, authoremail, scm, scmusername string) project.Project {
	return Golang{
		project.Info{
			License:     license,
			Author:      author,
			Authoremail: authoremail,
			Scm:         scm,
			ScmUserName: scmusername,
		},
	}
}

// Create creates a template folder structure for a golang project.
func (g Golang) Create(appname string) error {

	srcdir := appname

	if project.InitIfGitExist(srcdir) {
		log.Printf("Created repo: [%s] %s", project.GreenText(project.SignSuccess), srcdir)
	} else {
		// if git is not present, create basedir in the beginning, yourself
		project.MustCreateDir(srcdir)
	}

	makefilePath := path.Join(srcdir, "Makefile")
	readmePath := path.Join(srcdir, "README.md")
	testShellPath := path.Join(srcdir, "test.sh")
	gitignorePath := path.Join(srcdir, ".gitignore")
	mainPath := path.Join(srcdir, "main.go")
	licensePath := path.Join(srcdir, "LICENSE")

	data := project.Dict{
		"appname":     appname,
		"author":      g.Author,
		"year":        strconv.Itoa(time.Now().Year()),
		"license":     g.License,
		"authoremail": g.Authoremail,
		"scm":         g.Scm,
		"scmusername": g.ScmUserName,
	}
	parse := func(tpl string) string {
		return project.ParseTemplateString(tpl, t, data)
	}

	pathToContent := make(project.Dict)
	pathToContent[makefilePath] = parse(makefileText)
	pathToContent[readmePath] = parse(readmeText)
	pathToContent[testShellPath] = testShellText
	pathToContent[gitignorePath] = gitignoreText
	pathToContent[mainPath] = mainText
	pathToContent[licensePath] = parse(licenses.LicenseMap[g.License])

	for path, content := range pathToContent {
		err := project.WriteToFile(path, content)
		if err != nil {
			log.Printf("Failed file: [%s] %s", project.RedText(project.SignSuccess), path)
			return fmt.Errorf("Unable to write file: %s", path)
		}
		log.Printf("Created file: [%s] %s", project.GreenText(project.SignSuccess), path)
		time.Sleep(time.Millisecond * 100)
	}

	return nil
}
