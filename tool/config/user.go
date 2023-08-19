package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type UserConfig struct {
	Debug bool `yaml:"debug"`
	User  struct {
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"user"`
}

var UserConfigInstance *UserConfig

func GetUserConfigInstance() *UserConfig {
	if UserConfigInstance == nil {
		UserConfigInstance = &UserConfig{}
		UserConfigInstance.ReadConfig()
	}
	return UserConfigInstance
}

var UserConfigPath string = "config/user.yml"

func (c *UserConfig) ReadConfig() {

	configPath := UserConfigPath
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

func (c *UserConfig) IsDebug() bool {
	return GetUserConfigInstance().Debug
}
