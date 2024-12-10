package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

type envs struct {
	ALLOWED_ORIGINS string `json:"ALLOWED_ORIGINS" env:"ALLOWED_ORIGINS" default:""`
	LOG_LEVEL       string `json:"LOG_LEVEL" env:"LOG_LEVEL" default:"info"`
}

var (
	// pointer to access all envs of the project
	Envs = &envs{}
	// read rev.txt file to get application version string
	AppVersion string
)

func init() {
	LoadEnvs()
	version, err := os.ReadFile("rev.txt")
	if err != nil {
		version = []byte("development mode, no rev.txt")
	}
	AppVersion = string(version)
}

// LoadEnvs to reset envs in test context
func LoadEnvs() {
	envconfig.Process("", Envs)
}
