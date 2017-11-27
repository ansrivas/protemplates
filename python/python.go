package python

import (
	"fmt"
	"path"
	"strings"

	"github.com/ansrivas/protemplates/project"
)

// Python struct is responsible for creating golang projects
type Python struct{}

// New creates a new implementation for python which is later used to create a project.
func New(projectName string) project.Project {
	impl := Python{}
	return impl
}

// Create creates a template folder structure for a python project.
func (p Python) Create(appname string) error {

	appWithHyphen := strings.Replace(appname, "_", "-", -1)
	appWithUnderScore := strings.Replace(appname, "-", "_", -1)

	basedir := appname

	// Pip has some issues with project module containing `=`, so change it to `_`
	appdir := path.Join(appname, appWithUnderScore)
	testdir := path.Join(appname, "tests")

	dirs := []string{basedir, appdir, testdir}
	for _, dir := range dirs {
		project.MustCreateDir(dir)
	}

	pathToContent := make(map[string]string)

	//--------------------------------------------------------
	setuppyPath := path.Join(basedir, "setup.py")
	setupcfgPath := path.Join(basedir, "setup.cfg")
	gitignorePath := path.Join(basedir, ".gitignore")
	makefilePath := path.Join(basedir, "Makefile")
	requirementsPath := path.Join(basedir, "requirements.txt")
	readmePath := path.Join(basedir, "README.md")
	manifestPath := path.Join(basedir, "MANIFEST.in")
	devEnvYamlPath := path.Join(basedir, "dev_environment.yml")
	//--------------------------------------------------------

	//--------------------------------------------------------
	initpyPath := path.Join(appdir, "__init__.py")
	//--------------------------------------------------------

	//--------------------------------------------------------
	conftestPath := path.Join(testdir, "conftest.py")
	testfilePath := path.Join(testdir, "test_example_test.py")
	//--------------------------------------------------------

	// Pip has some issues with project module containing `=`, so change it to `_`
	pathToContent[setuppyPath] = fmt.Sprintf(setupyText, appWithUnderScore, appWithHyphen)

	//--------------------------------------------------------
	pathToContent[setupcfgPath] = fmt.Sprintf(setupCfgText, appWithUnderScore)
	pathToContent[gitignorePath] = gitignoreText
	pathToContent[conftestPath] = conftestText
	pathToContent[testfilePath] = testfileText
	pathToContent[initpyPath] = initpyText
	pathToContent[makefilePath] = makefileText
	pathToContent[requirementsPath] = requirementsText
	pathToContent[readmePath] = fmt.Sprintf(readmeText, appWithHyphen, appname, appname)
	pathToContent[manifestPath] = manifestText
	pathToContent[devEnvYamlPath] = fmt.Sprintf(devEnvYamlText, appWithHyphen)
	//--------------------------------------------------------

	for path, content := range pathToContent {
		err := project.WriteToFile(path, content)
		if err != nil {
			return fmt.Errorf("Unable to write file: %s", path)
		}
	}

	return nil
}
