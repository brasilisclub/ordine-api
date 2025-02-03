package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

type envs struct {
	ALLOWED_ORIGINS   string `json:"ALLOWED_ORIGINS" env:"ALLOWED_ORIGINS" default:""`
	DATABASE_HOST     string `json:"DATABASE_HOST" env:"DATABASE_HOST" default:"database"`
	DATABASE_PASSWORD string `json:"DATABASE_PASSWORD" env:"DATABASE_PASSWORD" default:"root"`
	DATABASE_PORT     string `json:"DATABASE_PORT" env:"DATABASE_PORT" default:"3306"`
	DATABASE_USER     string `json:"DATABASE_USER" env:"DATABASE_USER" default:"root"`
	ENVIRONMENT       string `json:"ENVIRONMENT" env:"ENVIRONMENT" default:"local"`
	JWT_SECRET_TOKEN  string `json:"JWT_SECRET_TOKEN" env:"JWT_SECRET_TOKEN" default:"secret"`
	LOG_LEVEL         string `json:"LOG_LEVEL" env:"LOG_LEVEL" default:"info"`
	PORT              string `json:"PORTS" env:"PORTS" default:"8080"`
}

var (
	// pointer to access all envs of the project
	Envs = &envs{}
	// read rev.txt file to get application version string
	AppVersion string
)

const Testing = "testing"

func IsApplicationInTestMode() bool {
	return Envs.ENVIRONMENT == Testing
}

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
