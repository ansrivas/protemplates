package main

import (
	"fmt"
	"os"
)

// This variable gets populated from the makefile or git tags automatically
var version = "undefined"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
