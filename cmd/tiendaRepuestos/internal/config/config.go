package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

//DBConfig ...
type DBConfig struct {
	Driver string `yaml:"driver"`
}

//Config ...
type Config struct {
	DB      DBConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

//LoadConfig ...
func LoadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c = &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
