package project

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	assert := assert.New(t)

	// Making this assumption because I have it + travis will have it
	expected := true

	tesdir := "./testdir"

	actual := InitIfGitExist(tesdir)
	defer os.RemoveAll(tesdir)

	assert.Equal(expected, actual, "Git exists, but failed to initialize directory")
}
