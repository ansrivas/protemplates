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

package project

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// CreateDir create a directory with a given name, only if it doesn't exist.
func CreateDir(dirname string) error {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		return os.Mkdir(dirname, os.ModePerm)
	}
	return fmt.Errorf("Unable to create directory %s", dirname)
}

// MustCreateDir creates a directory with given name/path. Panics if the directory already exists.
func MustCreateDir(dirpath string) {
	err := CreateDir(dirpath)
	if err != nil {
		panic(fmt.Errorf("Should have created directory: %s. Apparently it already exists", dirpath))
	}
}

// WriteToFile writes a given `data` string to a `filepath`.
func WriteToFile(filepath string, data string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}

// ParseTemplateString parses a string-template, fills in data and returns corresponding string
func ParseTemplateString(tpl string, t *template.Template, data interface{}) string {
	var b bytes.Buffer
	t, _ = t.Parse(tpl)
	t.Execute(&b, data)
	return b.String()
}

// SanitizeInput accepts a language string, checks if its allowed.
// If its allowed, returns a lowercase representation else error.
func SanitizeInput(input string) string {
	return strings.ToLower(input)
}

// Create accepts anything of interface type Project and calls underlying Create method.
func Create(p Project, appname string) error {
	return p.Create(appname)
}
