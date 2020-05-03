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

var configFileText = `
package config

import (
        "os"

        "github.com/joho/godotenv"
        "github.com/kelseyhightower/envconfig"
        "github.com/rs/zerolog/log"
)

const (
        configPrefix = "{{.appWithUnderScoreCapitalized}}"

        // EnvConfigPath represents the environment variable name which
        // should be read in case environment file needs to be read from some
        // user-defined location
        EnvConfigPath = "{{.appWithUnderScoreCapitalized}}_CONFIG_ENV_PATH"
)

// Config is the base struct which contains all the configuration for the
// application
type Config struct {

		// SomeUserName ....
        SomeUserName string ` + "`" + "envconfig:\"SOME_USER_NAME\" required:\"true\"" + "`" +
	`
        // SomeUserPass ...
        SomeUserPass string ` + "`" + "envconfig:\"SOME_USER_PASS\" default:\"some-default-value\"" + "`" +
	`
}

// LoadEnv will try to load .env file from the directory
// where it is currently running from, unless explicitly given
func LoadEnv() (Config, error) {

        pathToEnv := os.Getenv(EnvConfigPath)
        if pathToEnv == "" {
                pathToEnv = ".env"
        }

        log.Info().Msgf("Now reading config file %s", pathToEnv)
        var c Config
        err := godotenv.Load(pathToEnv)
        if err != nil {
                return c, err
        }

        err = envconfig.Process(configPrefix, &c)
        return c, err
}

`

var configTestFileText = `
package config

import (
        "os"
        "testing"

        "github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
        assert := assert.New(t)

        os.Setenv(EnvConfigPath, "env.test")
        expectedConfig := Config{
                SomeUserName:    "some-user-name",
                SomeUserPass:    "some-default-pass",
        }
        actualConfig, err := LoadEnv()
        assert.Nil(err, "Error in loading the config file")
        assert.Equal(expectedConfig, actualConfig, "Expected config is not equal to loaded config")
}

func TestLoadEnvFailure(t *testing.T) {
        assert := assert.New(t)
        os.Setenv(EnvConfigPath, "non.existing.env.test")
        _, err := LoadEnv()
        assert.NotNil(err, "Should not have read the configuration")
}
`

var envFileText = `
{{.appWithUnderScoreCapitalized}}_SOME_USER_NAME="some-user-name"
`
