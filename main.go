package main

import (
	"fmt"
	"os"
)

// Version gets populated from the makefile or git tags automatically
var Version = "undefined"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
