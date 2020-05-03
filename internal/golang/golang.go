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

package golang

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/ansrivas/protemplates/internal/licenses"
	"github.com/ansrivas/protemplates/internal/project"
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
	appWithUnderScoreCapitalized := strings.ToUpper(strings.Replace(appname, "-", "_", -1))

	scriptsDir := path.Join(appname, "scripts")
	internalDir := path.Join(appname, "internal")
	configDir := path.Join(internalDir, "config")
	if project.InitIfGitExist(srcdir) {
		log.Printf("Created repo: [%s] %s", project.GreenText(project.SignSuccess), srcdir)
	} else {
		// if git is not present, create basedir in the beginning, yourself
		project.MustCreateDir(srcdir)
	}
	project.MustCreateDir(scriptsDir)
	project.MustCreateDir(internalDir)
	project.MustCreateDir(configDir)

	// initialize go.mod
	normalize := func(input string) string {
		output := strings.TrimSuffix(input, "/")
		output = strings.TrimPrefix(input, "/")
		return output
	}

	goAppBasePath := fmt.Sprintf("%s/%s/%s", g.Scm, normalize(g.ScmUserName), normalize(appname))
	if !InitGoMod(srcdir, goAppBasePath) {
		log.Printf("Failed to initialize go.mod. Proceeding.")
	}

	configFilePath := path.Join(configDir, "config.go")
	configTestFilePath := path.Join(configDir, "config_test.go")
	configEnvFilePath := path.Join(configDir, "env.test")

	makefilePath := path.Join(srcdir, "Makefile")
	readmePath := path.Join(srcdir, "README.md")
	testShellPath := path.Join(scriptsDir, "test.sh")
	gitignorePath := path.Join(srcdir, ".gitignore")
	mainPath := path.Join(srcdir, "main.go")
	mainTestPath := path.Join(srcdir, "main_test.go")
	licensePath := path.Join(srcdir, "LICENSE")
	changelogPath := path.Join(srcdir, "CHANGELOG.md")

	data := project.Dict{
		"appname":                      appname,
		"appWithUnderScoreCapitalized": appWithUnderScoreCapitalized,
		"goAppBasePath":                goAppBasePath,
		"author":                       g.Author,
		"year":                         strconv.Itoa(time.Now().Year()),
		"license":                      g.License,
		"authoremail":                  g.Authoremail,
		"scm":                          g.Scm,
		"scmusername":                  g.ScmUserName,
	}
	parse := func(tpl string) string {
		return project.ParseTemplateString(tpl, t, data)
	}

	pathToContent := make(project.Dict)

	pathToContent[configFilePath] = parse(configFileText)
	pathToContent[configTestFilePath] = parse(configTestFileText)
	pathToContent[configEnvFilePath] = parse(envFileText)

	pathToContent[makefilePath] = parse(makefileText)
	pathToContent[readmePath] = parse(readmeText)
	pathToContent[testShellPath] = testShellText
	pathToContent[gitignorePath] = gitignoreText
	pathToContent[mainPath] = parse(mainText)
	pathToContent[mainTestPath] = mainTestText
	pathToContent[licensePath] = parse(licenses.LicenseMap[g.License])
	pathToContent[changelogPath] = changelogText

	for path, content := range pathToContent {
		err := project.WriteToFile(path, content)
		if err != nil {
			log.Printf("Failed file: [%s] %s", project.RedText(project.SignSuccess), path)
			return fmt.Errorf("Unable to write file: %s", path)
		}
		log.Printf("Created file: [%s] %s", project.GreenText(project.SignSuccess), path)
		time.Sleep(time.Millisecond * 100)
	}
	project.InitialCommit(srcdir)
	return nil
}
