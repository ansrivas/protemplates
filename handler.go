package main

import (
	"fmt"
	"os"

	"github.com/ansrivas/protemplates/golang"
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

	var impl project.Project
	switch lang {
	case "python":
		impl = python.New(projectName)
	case "go", "golang":
		impl = golang.New(projectName)
	default:
		fmt.Printf("\033[31mException: Language %s is currently not supported.\033[39m\n\n", language)
		os.Exit(1)
	}
	mustCreateProject(impl, projectName)
}
