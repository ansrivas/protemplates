package main

import (
	"fmt"
	"os"
)

var version = "undefined"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
