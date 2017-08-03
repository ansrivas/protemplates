package main

import (
	"fmt"
	"os"

	"github.com/ansrivas/protemplates/internal"
	"github.com/ansrivas/protemplates/python"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "protemplates"}
)

func init() {
	rootCmd.AddCommand(cmdCreate)
	rootCmd.AddCommand(cmdVersion)
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of protemplates",
	Long:  `All softwares have versions. This is protemplate's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Protemplates: ", version)
	},
}

var cmdCreate = &cobra.Command{
	Use:     "create [creates a project template for a given language]",
	Short:   "Creates a project template for a given language.",
	Long:    "Creates a project template for a given language",
	Example: "protemplates create python",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			curcmd, _, _ := rootCmd.Find([]string{"create"})
			fmt.Printf("\033[31m%s\033[39m\n\n", curcmd.UsageString())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		handleProjectCreation(args[0])
	},
}

func handleProjectCreation(language string) {
	lang := internal.SanitizeInput(language)

	var implementation internal.Project

	switch lang {
	case "python":
		implementation = python.Python{}
	default:
		fmt.Printf("\033[31mException: %s is currently not supported.\033[39m\n\n", language)
		os.Exit(1)
	}

	var projectName string
	fmt.Println("Please enter a desired project name:")
	fmt.Scanf("%s", &projectName)
	err := internal.Create(implementation, projectName)
	if err != nil {
		panic(fmt.Sprintf("Unable to create the project: %s", projectName))
	}
	fmt.Printf("Successfully created python project %s in current directory\n", projectName)
}
