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
