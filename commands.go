package main

import (
	"fmt"
	"os"

	"github.com/ansrivas/ptemplate/python"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "ptemplate"}
)

func init() {

	rootCmd.AddCommand(cmdCreate)

}

var cmdCreate = &cobra.Command{
	Use:     "create [creates a project template for a given language]",
	Short:   "Creates a project template for a given language.",
	Long:    "Creates a project template for a given language",
	Example: "ptemplate create python",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			curcmd, _, _ := rootCmd.Find([]string{"create"})
			fmt.Printf("\033[31m%s\033[39m\n\n", curcmd.UsageString())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		fmt.Println("Please enter a desired project name:")
		fmt.Scanf("%s", &projectName)
		err := python.Create(projectName)
		if err != nil {
			fmt.Printf("Successfully created python project %s in current directory", projectName)
		}
	},
}
