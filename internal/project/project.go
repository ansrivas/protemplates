package project

// Project interface to support creation of different projects.
type Project interface {

	// Create will be implemented by structs which want to create
	// a directory structure for a given language eg. Python.
	Create(appname string) error
}

// Dict is a dict of string keys to string values
type Dict map[string]string
