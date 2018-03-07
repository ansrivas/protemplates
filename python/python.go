package python

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/ansrivas/protemplates/licenses"
	"github.com/ansrivas/protemplates/project"
)

// Python struct is responsible for creating python projects
type Python struct {
	project.Info
}

var t = template.New("python")

// New creates a new implementation for python which is later used to create a project.
func New(projectName, license, author, authoremail, scm, scmusername string) project.Project {

	return Python{
		project.Info{
			License:     license,
			Author:      author,
			Authoremail: authoremail,
			Scm:         scm,
			ScmUserName: scmusername,
		},
	}
}

// Create creates a template folder structure for a python project.
func (p Python) Create(appname string) error {

	appWithHyphen := strings.Replace(appname, "_", "-", -1)
	appWithUnderScore := strings.Replace(appname, "-", "_", -1)

	basedir := appname

	// Pip has some issues with project module containing `=`, so change it to `_`
	appdir := path.Join(appname, appWithUnderScore)
	testdir := path.Join(appname, "tests")
	examplesdir := path.Join(appname, "examples")

	dirs := []string{appdir, testdir, examplesdir}
	// creates the basedir and initializes an empty repo
	if project.InitIfGitExist(basedir) {
		log.Printf("Created repo: [%s] %s", project.GreenText(project.SignSuccess), basedir)
	} else {
		// if git is not present, create basedir in the beginning, yourself
		dirs = append([]string{basedir}, dirs...)
	}

	for _, dir := range dirs {
		project.MustCreateDir(dir)
	}

	setuppyPath := path.Join(basedir, "setup.py")
	setupcfgPath := path.Join(basedir, "setup.cfg")
	gitignorePath := path.Join(basedir, ".gitignore")
	makefilePath := path.Join(basedir, "Makefile")
	requirementsPath := path.Join(basedir, "requirements.txt")
	readmePath := path.Join(basedir, "README.md")
	manifestPath := path.Join(basedir, "MANIFEST.in")
	devEnvYamlPath := path.Join(basedir, "dev_environment.yml")
	activateEnvPath := path.Join(basedir, "activate-env.sh")
	travisYmlPath := path.Join(basedir, ".travis.yml")
	licensePath := path.Join(basedir, "LICENSE")

	initpyPath := path.Join(appdir, "__init__.py")

	conftestPath := path.Join(testdir, "conftest.py")
	testfilePath := path.Join(testdir, "test_init.py")

	examplesPath := path.Join(examplesdir, "simple.py")

	data := project.Dict{
		"appname":           appname,
		"appWithUnderScore": appWithUnderScore,
		"appWithHyphen":     appWithHyphen,
		"author":            p.Author,
		"year":              strconv.Itoa(time.Now().Year()),
		"license":           p.License,
		"authoremail":       p.Authoremail,
		"scm":               p.Scm,
		"scmusername":       p.ScmUserName,
	}
	parse := func(tpl string) string {
		return project.ParseTemplateString(tpl, t, data)
	}

	pathToContent := make(project.Dict)
	pathToContent[setuppyPath] = parse(setupyText)
	pathToContent[setupcfgPath] = parse(setupCfgText)
	pathToContent[gitignorePath] = gitignoreText
	pathToContent[conftestPath] = conftestText
	pathToContent[testfilePath] = parse(testfileText)
	pathToContent[initpyPath] = initpyText
	pathToContent[makefilePath] = makefileText
	pathToContent[requirementsPath] = requirementsText
	pathToContent[readmePath] = parse(readmeText())
	pathToContent[manifestPath] = manifestText
	pathToContent[devEnvYamlPath] = parse(devEnvYamlText)
	pathToContent[activateEnvPath] = parse(activateEnvText)
	pathToContent[travisYmlPath] = travisText
	pathToContent[examplesPath] = parse(examplesText)
	pathToContent[licensePath] = parse(licenses.LicenseMap[p.License])

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
