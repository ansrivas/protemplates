package internal

// Project interface to support creation of different projects.
type Project interface {

	// Create will be implemented by structs which want to create
	// a directory structure for a given language eg. Python.
	Create(appname string) error
}
