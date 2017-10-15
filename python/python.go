package python

import (
	"fmt"
	"path"
	"strings"

	"github.com/ansrivas/protemplates/internal"
)

// Python ...
type Python struct{}

// Create creates a template folder structure for a python project.
func (p Python) Create(appname string) error {
	basedir := appname
	appdir := path.Join(appname, appname)
	testdir := path.Join(appname, "tests")

	dirs := []string{basedir, appdir, testdir}
	for _, dir := range dirs {
		internal.MustCreateDir(dir)
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
	//--------------------------------------------------------

	//--------------------------------------------------------
	initpyPath := path.Join(appdir, "__init__.py")
	//--------------------------------------------------------

	//--------------------------------------------------------
	conftestPath := path.Join(testdir, "conftest.py")
	testfilePath := path.Join(testdir, "test_first.py")
	//--------------------------------------------------------

	pathToContent[setuppyPath] = fmt.Sprintf(setupyText, appname, appname)

	//--------------------------------------------------------
	pathToContent[setupcfgPath] = setupCfgText
	pathToContent[gitignorePath] = gitignoreText
	pathToContent[conftestPath] = conftestText
	pathToContent[testfilePath] = testfileText
	pathToContent[initpyPath] = initpyText
	pathToContent[makefilePath] = makefileText
	pathToContent[requirementsPath] = requirementsText
	pathToContent[readmePath] = fmt.Sprintf(readmeText, strings.Title(appname))
	pathToContent[manifestPath] = manifestText
	//--------------------------------------------------------

	for path, content := range pathToContent {
		err := internal.WriteToFile(path, content)
		if err != nil {
			return fmt.Errorf("Unable to write file: %s", path)
		}
	}

	return nil
}
