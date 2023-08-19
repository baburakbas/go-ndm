package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AssetConfig struct {
	Debug bool `yaml:"debug"`
	Asset struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"asset"`
}

var AssetConfigInstance *AssetConfig

func GetAssetConfigInstance() *AssetConfig {
	if AssetConfigInstance == nil {
		AssetConfigInstance = &AssetConfig{}
		AssetConfigInstance.ReadConfig()
	}
	return AssetConfigInstance
}

var AssetConfigPath string = "config/asset.yml"

func (c *AssetConfig) ReadConfig() {

	configPath := AssetConfigPath
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

func (c *AssetConfig) IsDebug() bool {
	return GetAssetConfigInstance().Debug
}
