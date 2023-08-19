package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Debug bool `yaml:"debug"`
	App   struct {
		FilePath   string `yaml:"filepath"`
		FilePrefix string `yaml:"fileprefix"`
		FileEndfix string `yaml:"fileendfix"`
	} `yaml:"app"`
}

var AppConfigInstance *AppConfig

func GetAppConfigInstance() *AppConfig {
	if AppConfigInstance == nil {
		AppConfigInstance = &AppConfig{}
		AppConfigInstance.ReadConfig()
	}
	return AppConfigInstance
}

var AppConfigPath string = "config/app.yml"

func (c *AppConfig) ReadConfig() {

	configPath := AppConfigPath
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

func (c *AppConfig) IsDebug() bool {
	return GetAppConfigInstance().Debug
}
