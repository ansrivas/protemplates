package project

// Info represents bunch of information associated with a project
type Info struct {
	// License represents a string like "MIT" or "Apache2"
	License string
	// Author represents the full name of author. For eg. Tova Lanre
	Author string
	// Authoremail represents the email of the author in case anyone needs to contact
	Authoremail string
	// Scm is source code management, for eg. github.com, gitlab.com, bitbucket.com
	Scm string
	// ScmUserName represents a username, for eg. github.com/spf13 => then spf13
	ScmUserName string
}
