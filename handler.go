package main

import (
	"fmt"
	"os"

	"github.com/ansrivas/protemplates/golang"
	"github.com/ansrivas/protemplates/licenses"
	"github.com/ansrivas/protemplates/project"
	"github.com/ansrivas/protemplates/python"
)

// mustCreateProject tries to create a project and panics if something fails.
func mustCreateProject(impl project.Project, projectName string) {
	err := project.Create(impl, projectName)
	if err != nil {
		panic(fmt.Sprintf("Unable to create the project: %s", projectName))
	}
	fmt.Printf("Successfully created the project %s in current directory\n", projectName)
}

func createProject(language string) {
	lang := project.SanitizeInput(language)

	var projectName string
	fmt.Println("Please enter a desired project name:")
	for {
		fmt.Scanf("%s", &projectName)
		if projectName != "" {
			break
		}
		fmt.Println("Project name can not be empty, please enter a valid project name:")
	}

	fmt.Println(`Please enter a license type.
[1] MIT
[2] Apache2`)

	var license string
	for {
		license = ""
		var licenseInput int
		fmt.Scanf("%d", &licenseInput)
		switch licenseInput {
		case 1:
			license = licenses.Mit
		case 2:
			license = licenses.Apache2
		default:
			fmt.Println("Licenses can be either MIT or Apache2 currently")
		}
		if license != "" {
			break
		}
		continue
	}

	fmt.Println(`Please enter author name:`)

	var author string
	for {
		fmt.Scanf("%s", &author)
		if author != "" {
			break
		}
		fmt.Println("Author name can not be empty, please enter a valid name:")
	}

	var impl project.Project
	switch lang {
	case "python":
		impl = python.New(projectName, license, author)
	case "go", "golang":
		impl = golang.New(projectName)
	default:
		fmt.Printf("\033[31mException: Language %s is currently not supported.\033[39m\n\n", language)
		os.Exit(1)
	}
	mustCreateProject(impl, projectName)
}
