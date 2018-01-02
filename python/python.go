package python

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/ansrivas/protemplates/licenses"
	"github.com/ansrivas/protemplates/project"
)

// Python struct is responsible for creating golang projects
type Python struct {
	license string
	author  string
}

// New creates a new implementation for python which is later used to create a project.
func New(projectName, license, author string) project.Project {
	impl := Python{license: license,
		author: author}
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
	travisYmlPath := path.Join(basedir, ".travis.yml")
	licensePath := path.Join(basedir, "LICENSE")
	//--------------------------------------------------------

	//--------------------------------------------------------
	initpyPath := path.Join(appdir, "__init__.py")
	//--------------------------------------------------------

	//--------------------------------------------------------
	conftestPath := path.Join(testdir, "conftest.py")
	testfilePath := path.Join(testdir, "test_init.py")
	//--------------------------------------------------------

	examplesPath := path.Join(examplesdir, "simple.py")

	//=================================================================

	pathToContent[setuppyPath] = fmt.Sprintf(setupyText, appWithUnderScore, appWithHyphen, p.author)

	//--------------------------------------------------------
	pathToContent[setupcfgPath] = fmt.Sprintf(setupCfgText, appWithUnderScore)
	pathToContent[gitignorePath] = gitignoreText
	pathToContent[conftestPath] = conftestText
	pathToContent[testfilePath] = fmt.Sprintf(testfileText, appWithUnderScore, appWithUnderScore)
	pathToContent[initpyPath] = initpyText
	pathToContent[makefilePath] = makefileText
	pathToContent[requirementsPath] = requirementsText
	pathToContent[readmePath] = fmt.Sprintf(readmeText, appWithHyphen, appname, appname, appname, appname, p.license)
	pathToContent[manifestPath] = manifestText
	pathToContent[devEnvYamlPath] = fmt.Sprintf(devEnvYamlText, appWithHyphen)
	pathToContent[travisYmlPath] = travisText
	pathToContent[licensePath] = fmt.Sprintf(licenses.LicenseMap[p.license], strconv.Itoa(time.Now().Year()), p.author)
	pathToContent[examplesPath] = fmt.Sprintf(examplesText, appWithUnderScore, appWithUnderScore)
	//--------------------------------------------------------

	for path, content := range pathToContent {
		err := project.WriteToFile(path, content)
		if err != nil {
			return fmt.Errorf("Unable to write file: %s", path)
		}
		log.Printf("Created file: [%s] %s", project.GreenText(project.SignSuccess), path)
		time.Sleep(time.Millisecond * 100)
	}

	return nil
}
