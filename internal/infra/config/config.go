package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	LogLevel string `yaml:"log_level"`
}

var (
	_, b, _, _ = runtime.Caller(0)
)

func InitConfiguration() Configuration {
	var path string
	if getEnvironment() == "live" {
		path = fmt.Sprintf("%s/application-%s.yaml", getConfigFolderPath(), getEnvironment())
	} else {
		root := filepath.Join(filepath.Dir(b), "../../../")
		path = fmt.Sprintf("%s/config/application-%s.yaml", root, getEnvironment())
	}

	cfg, err := load(path)

	if err != nil {
		panic(err)
	}
	return *cfg
}

func load(path string) (*Configuration, error) {
	var configuration *Configuration

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))

	err = yaml.Unmarshal(yamlFile, &configuration)
	if err != nil {
		return nil, err
	}

	return configuration, nil
}

func getConfigFolderPath() string {
	pwd, _ := os.Getwd()
	return pwd + "/config"
}

func getEnvironment() string {
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "local"
	}

	return environment
}
