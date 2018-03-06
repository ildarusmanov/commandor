package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// config data
type Config struct {
	secretKey string `yaml:"secret_key"`
	tokenType string `yaml:"token_type"`
}

// config constructor
func CreateNewConfig() *Config {
	return &Config{}
}

// load config from YAML file
// returns *Config
func CreateFromYAML(yamlFilePath string) *Config {
	c := CreateNewConfig()

	yamlData, err := ioutil.ReadFile(yamlFilePath)

	if err != nil {
		return nil, err
	}

	if yaml.Unmarshal([]byte(c), yamlData) != nil {
		return nil, err
	}

	return c, nil
}

// secret key getter
func (c *Config) GetSecretKey() string {
	return c.secretKey
}

// token type getter
func (c *Config) GetTokenType() string {
	return c.tokenType
}
