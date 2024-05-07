package config

import "os"

const EnvKey = "TOY_ENV"
const Dev = "development"
const Test = "test"
const Production = "production"

type AppConfig struct {
	Env     string
	Db      *DbConfig
	Setting *Setting
}

type DbConfig struct {
	Url      string `yaml:"url"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Setting struct{}

func GetConfig() (AppConfig, error) {
	return AppConfig{}, nil
}

func getDbConfig() {
	env := GetEnv()

}

func GetEnv() string {
	env := os.Getenv(EnvKey)
	if env == "" {
		env = Dev
	}
	return env
}
