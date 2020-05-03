//
// Copyright (c) 2020 Ankur Srivastava
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package golang

var mainText = `package main

import (
        "flag"
        "fmt"
        "os/signal"
        "syscall"

        "os"
        figure "github.com/common-nighthawk/go-figure"
        config "{{.goAppBasePath}}/internal/config"
		
        "github.com/rs/zerolog"
        "github.com/rs/zerolog/log"

)

var (
	// BuildTime gets populated during the build proces
	BuildTime = ""

	//Version gets populated during the build process
	Version = ""
)


// setupConfigOrFatal loads all the variables from the environment variable.
// At this point everything is read as a Key,Value in a map[string]string
func setupConfigOrFatal() config.Config {
        conf, err := config.LoadEnv()
        if err != nil {

                log.Fatal().Msgf("Failed to parse the environment variable. Error %s", err.Error())
        }
        return conf
}

func printBanner() {
        myFigure := figure.NewFigure("{{.appname}}", "", true)
        myFigure.Print()
}


// setupLogger will setup the zap json logging interface
// if the --debug flag is passed, level will be debug
func setupLogger(debug bool) {

        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
        zerolog.SetGlobalLevel(zerolog.InfoLevel)
        if debug {
                zerolog.SetGlobalLevel(zerolog.DebugLevel)
        }

}

// printVersionInfo just prints the build time and git commit/tag used
// for this build
func printVersionInfo(version bool) {
        if version {
                fmt.Printf("Version  : %s\nBuildTime: %s\n", Version, BuildTime)
                os.Exit(0)
        }
}

func helloWorld() string{
	return "Hello World"
}

func main() {
		debug := flag.Bool("debug", false, "Set the log level to debug")
        version := flag.Bool("version", false, "Display the BuildTime and Version of this binary")
        flag.Parse()
		
		printVersionInfo(*version)
        setupLogger(*debug)

        errc := make(chan error)
        go func() {
                c := make(chan os.Signal)
                signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
                errc <- fmt.Errorf("%s", <-c)
		}()

        log.Info().Msgf("Exiting server. Message: %v", <-errc)
}
`

var mainTestText = `package main

import(
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_helloWorld(t *testing.T) {
	assert := assert.New(t)
	expected:= "Hello World"
	assert.Equal( expected, helloWorld(), "Failed test HelloWorld")
}
`
